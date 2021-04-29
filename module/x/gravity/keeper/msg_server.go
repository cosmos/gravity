package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/gravity-bridge/module/x/gravity/types"
)

var _ types.MsgServer = &Keeper{}

// SetDelegateKey implements MsgServer.SetDelegateKey. The
func (k Keeper) SetDelegateKey(c context.Context, msg *types.MsgDelegateKey) (*types.MsgDelegateKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// NOTE: address checked on msg validation
	validatorAddr, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	orchestratorAddr, _ := sdk.AccAddressFromBech32(msg.OrchestratorAddress)
	ethereumAddr := common.HexToAddress(msg.EthAddress)

	// ensure that the validator exists
	if k.stakingKeeper.Validator(ctx, validatorAddr) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, validatorAddr.String())
	}

	// TODO: consider impact of maliciously setting duplicate delegate
	// addresses since no signatures from the private keys of these addresses
	// are required for this message it could be sent in a hostile way.

	// set the orchestrator address (map[orch]val)
	k.SetOrchestratorValidator(ctx, validatorAddr, orchestratorAddr)
	// set the ethereum address (map[val]eth)
	k.SetEthAddress(ctx, validatorAddr, ethereumAddr)
	// set third index (map[eth]orch)
	k.SetEthOrchAddress(ctx, ethereumAddr, orchestratorAddr)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySetOperatorAddr, msg.OrchestratorAddress),
		),
	)

	return &types.MsgDelegateKeyResponse{}, nil
}

func (k Keeper) SubmitEvent(c context.Context, msg *types.MsgSubmitEvent) (*types.MsgSubmitEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// NOTE: error checked on msg validate basic
	orchestratorAddr, _ := sdk.AccAddressFromBech32(msg.Signer)

	event, err := types.UnpackEvent(msg.Event)
	if err != nil {
		return nil, err
	}

	if err := k.HandleEthEvent(ctx, event, orchestratorAddr); err != nil {
		return nil, err
	}

	return &types.MsgSubmitEventResponse{}, nil
}

// FIXME:
func (k Keeper) SubmitConfirm(c context.Context, msg *types.MsgSubmitConfirm) (*types.MsgSubmitConfirmResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// NOTE: error checked on msg validate basic
	orchestratorAddr, _ := sdk.AccAddressFromBech32(msg.Signer)
	validatorAddr := k.GetOrchestratorValidator(ctx, orchestratorAddr)
	if validatorAddr == nil {
		return nil, sdkerrors.Wrapf(stakingtypes.ErrNoValidatorFound, "orchestrator address %s", orchestratorAddr)
	}

	confirm, err := types.UnpackConfirm(msg.Confirm)
	if err != nil {
		return nil, err
	}

	if err := k.ConfirmEvent(ctx, confirm, validatorAddr); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyValidator, validatorAddr.String()),
			sdk.NewAttribute(types.AttributeKeyConfirmType, confirm.GetType()),
		),
	)

	return &types.MsgSubmitConfirmResponse{}, nil
}

// RequestBatch handles MsgRequestBatch
func (k Keeper) RequestBatch(c context.Context, msg *types.MsgRequestBatch) (*types.MsgRequestBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// orchestratorAddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)

	// FIXME: update logic

	var (
		tokenContract common.Address
		found         bool
	)

	if types.IsEthereumERC20Token(msg.Denom) {
		tokenContractHex := types.GravityDenomToERC20Contract(msg.Denom)
		tokenContract = common.HexToAddress(tokenContractHex)
	} else {
		// get contract from store
		tokenContract, found = k.GetERC20ContractFromCoinDenom(ctx, msg.Denom)
		if !found {
			// TODO: what if there is no corresponding contract? will it be "generated" on ethereum
			// upon receiving?
			// FIXME: this will fail if the cosmos tokens are relayed for the first time and they don't have a counterparty contract
			// Fede: Also there's the question of how do we handle IBC denominations from a security perspective. Do we assign them the same
			// contract? My guess is that each new contract assigned to a cosmos coin should be approved by governance
			return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "denom %s", msg.Denom)
		}
	}

	_, err := k.CreateBatchTx(ctx, tokenContract)
	if err != nil {
		return nil, err
	}

	// TODO: later make sure that Demon matches a list of tokens already
	// in the bridge to send

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
		),
	)

	return &types.MsgRequestBatchResponse{}, nil
}

// Transfer handles MsgTransfer
func (k Keeper) Transfer(c context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// NOTE: errors checked on msg validation
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	ethereumAddr := common.HexToAddress(msg.EthRecipient)

	txID, err := k.AddTransferToOutgoingPool(ctx, sender, ethereumAddr, msg.Amount, msg.BridgeFee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyEthRecipient, msg.EthRecipient),
			sdk.NewAttribute(types.AttributeKeyTxID, txID.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Amount.Denom),
		),
	)

	return &types.MsgTransferResponse{
		TxId: txID,
	}, nil
}

// CancelTransfer
func (k Keeper) CancelTransfer(c context.Context, msg *types.MsgCancelTransfer) (*types.MsgCancelTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	sender, _ := sdk.AccAddressFromBech32(msg.Sender)

	if err := k.RemoveFromOutgoingPoolAndRefund(ctx, msg.TxId, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyTxID, msg.TxId.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)

	return &types.MsgCancelTransferResponse{}, nil
}
