package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/althea-net/peggy/module/x/peggy/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) DelegateKey(c context.Context, msg *types.MsgDelegateKeys) (*types.MsgDelegateKeysResponse, error) {
	// ensure that this passes validation
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	val, _ := sdk.ValAddressFromBech32(msg.Validator)
	orch, _ := sdk.AccAddressFromBech32(msg.Orchestrator)

	// ensure that the validator exists
	if k.Keeper.StakingKeeper.Validator(ctx, val) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}

	// TODO consider impact of maliciously setting duplicate delegate
	// addresses since no signatures from the private keys of these addresses
	// are required for this message it could be sent in a hostile way.

	// set the orchestrator address
	k.SetOrchestratorValidator(ctx, val, orch)
	// set the ethereum address
	k.SetEthAddress(ctx, val, msg.EthAddress)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySetOperatorAddr, orch.String()),
		),
	)

	return &types.MsgSetOrchestratorAddressResponse{}, nil

}

// ValsetConfirm handles MsgValsetConfirm
// TODO: check msgValsetConfirm to have an Orchestrator field instead of a Validator field
// func (k msgServer) ValsetConfirm(c context.Context, msg *types.MsgValsetConfirm) (*types.MsgValsetConfirmResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)
// 	valset := k.GetValset(ctx, msg.Nonce)
// 	if valset == nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find valset")
// 	}

// 	peggyID := k.GetPeggyID(ctx)
// 	checkpoint := valset.GetCheckpoint(peggyID)

// 	sigBytes, err := hex.DecodeString(msg.Signature)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
// 	}

// 	valaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
// 	validator := k.GetOrchestratorValidator(ctx, valaddr)
// 	if validator == nil {
// 		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(valaddr))
// 		if sval == nil {
// 			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
// 		}
// 		validator = sval.GetOperator()
// 	}

// 	ethAddress := k.GetEthAddress(ctx, validator)
// 	if ethAddress == "" {
// 		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
// 	}

// 	if err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress); err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with peggy-id %s with checkpoint %s found %s", ethAddress, peggyID, hex.EncodeToString(checkpoint), msg.Signature))
// 	}

// 	// persist signature
// 	if k.GetValsetConfirm(ctx, msg.Nonce, valaddr) != nil {
// 		return nil, sdkerrors.Wrap(types.ErrDuplicate, "signature duplicate")
// 	}
// 	key := k.SetValsetConfirm(ctx, *msg)

// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
// 			sdk.NewAttribute(types.AttributeKeyValsetConfirmKey, string(key)),
// 		),
// 	)

// 	return &types.MsgValsetConfirmResponse{}, nil
// }

// SendToEth handles MsgSendToEth
func (k msgServer) SendToEth(c context.Context, msg *types.MsgSendToEth) (*types.MsgSendToEthResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	txID, err := k.AddToOutgoingPool(ctx, sender, msg.EthDest, msg.Amount, msg.BridgeFee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyOutgoingTXID, fmt.Sprint(txID)),
		),
	)

	return &types.MsgSendToEthResponse{}, nil
}

// RequestBatch handles MsgRequestBatch
func (k msgServer) RequestBatch(c context.Context, msg *types.MsgRequestBatch) (*types.MsgRequestBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the denom is a peggy coin, if not, check if there is a deployed ERC20 representing it.
	// If not, error out
	_, tokenContract, err := k.DenomToERC20Lookup(ctx, msg.Denom)
	if err != nil {
		return nil, err
	}

	batchID, err := k.BuildOutgoingTXBatch(ctx, tokenContract, OutgoingTxBatchSize)
	if err != nil {
		return nil, err
	}

	// question: is this right? If i can delegate my voting power to a different key then this would fail each time i call it
	valaddr, _ := sdk.AccAddressFromBech32(msg.Validator)
	validator := k.GetOrchestratorValidator(ctx, valaddr)
	if validator == nil {
		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(valaddr))
		if sval == nil {
			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
		}
		validator = sval.GetOperator()
	}

	// TODO later make sure that Demon matches a list of tokens already
	// in the bridge to send

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyBatchNonce, fmt.Sprint(batchID.BatchNonce)),
		),
	)

	return &types.MsgRequestBatchResponse{}, nil
}

// ConfirmBatch handles MsgConfirmBatch
func (k msgServer) SubmitConfirm(c context.Context, msg *types.MsgSubmitConfirm) (*types.MsgSubmitConfirmResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// fetch the outgoing batch given the nonce
	batch := k.GetOutgoingTXBatch(ctx, msg.TokenContract, msg.Nonce)
	if batch == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find batch")
	}

	peggyID := k.GetPeggyID(ctx)
	checkpoint, err := batch.GetCheckpoint(peggyID)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "checkpoint generation")
	}

	sigBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
	}

	valaddr, _ := sdk.AccAddressFromBech32(msg.Signer)
	validator := k.GetOrchestratorValidator(ctx, valaddr)
	if validator == nil {
		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(valaddr))
		if sval == nil {
			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
		}
		validator = sval.GetOperator()
	}

	ethAddress := k.GetEthAddress(ctx, validator)
	if ethAddress == "" {
		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
	}

	err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with peggy-id %s with checkpoint %s found %s", ethAddress, peggyID, hex.EncodeToString(checkpoint), msg.Signature))
	}

	// check if we already have this confirm
	if k.GetBatchConfirm(ctx, msg.Nonce, msg.TokenContract, valaddr) != nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicate, "duplicate signature")
	}
	key := k.SetBatchConfirm(ctx, msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyBatchConfirmKey, string(key)),
		),
	)

	return &MsgSubmitConfirmResponse{}, nil
}

// // ConfirmLogicCall handles MsgConfirmLogicCall
// func (k msgServer) ConfirmLogicCall(c context.Context, msg *types.MsgConfirmLogicCall) (*types.MsgConfirmLogicCallResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)
// 	invalidationIdBytes, err := hex.DecodeString(msg.InvalidationId)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "invalidation id encoding")
// 	}

// 	// fetch the outgoing logic given the nonce
// 	logic := k.GetOutgoingLogicCall(ctx, invalidationIdBytes, msg.InvalidationNonce)
// 	if logic == nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "couldn't find logic")
// 	}

// 	peggyID := k.GetPeggyID(ctx)
// 	checkpoint, err := logic.GetCheckpoint(peggyID)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "checkpoint generation")
// 	}

// 	sigBytes, err := hex.DecodeString(msg.Signature)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, "signature decoding")
// 	}

// 	valaddr, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
// 	validator := k.GetOrchestratorValidator(ctx, valaddr)
// 	if validator == nil {
// 		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(valaddr))
// 		if sval == nil {
// 			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
// 		}
// 		validator = sval.GetOperator()
// 	}

// 	ethAddress := k.GetEthAddress(ctx, validator)
// 	if ethAddress == "" {
// 		return nil, sdkerrors.Wrap(types.ErrEmpty, "eth address")
// 	}

// 	err = types.ValidateEthereumSignature(checkpoint, sigBytes, ethAddress)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(types.ErrInvalid, fmt.Sprintf("signature verification failed expected sig by %s with peggy-id %s with checkpoint %s found %s", ethAddress, peggyID, hex.EncodeToString(checkpoint), msg.Signature))
// 	}

// 	// check if we already have this confirm
// 	if k.GetLogicCallConfirm(ctx, invalidationIdBytes, msg.InvalidationNonce, valaddr) != nil {
// 		return nil, sdkerrors.Wrap(types.ErrDuplicate, "duplicate signature")
// 	}

// 	k.SetLogicCallConfirm(ctx, msg)

// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
// 		),
// 	)

// 	return nil, nil
// }

// DepositClaim handles MsgDepositClaim
// TODO it is possible to submit an old msgDepositClaim (old defined as covering an event nonce that has already been
// executed aka 'observed' and had it's slashing window expire) that will never be cleaned up in the endblocker. This
// should not be a security risk as 'old' events can never execute but it does store spam in the chain.
// func (k msgServer) DepositClaim(c context.Context, msg *types.MsgDepositClaim) (*types.MsgDepositClaimResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	orch, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
// 	validator := k.GetOrchestratorValidator(ctx, orch)
// 	if validator == nil {
// 		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(orch))
// 		if sval == nil {
// 			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
// 		}
// 		validator = sval.GetOperator()
// 	}

// 	// return an error if the validator isn't in the active set
// 	val := k.StakingKeeper.Validator(ctx, validator)
// 	if val == nil || !val.IsBonded() {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in active set")
// 	}

// 	any, err := codectypes.NewAnyWithValue(msg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Add the claim to the store
// 	_, err = k.Attest(ctx, msg, any)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(err, "create attestation")
// 	}

// 	// Emit the handle message event
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
// 			// TODO: maybe return something better here? is this the right string representation?
// 			sdk.NewAttribute(types.AttributeKeyAttestationID, string(types.GetAttestationKey(msg.EventNonce, msg.ClaimHash()))),
// 		),
// 	)

// 	return &types.MsgDepositClaimResponse{}, nil
// }

// WithdrawClaim handles MsgWithdrawClaim
// TODO it is possible to submit an old msgWithdrawClaim (old defined as covering an event nonce that has already been
// executed aka 'observed' and had it's slashing window expire) that will never be cleaned up in the endblocker. This
// should not be a security risk as 'old' events can never execute but it does store spam in the chain.
// func (k msgServer) WithdrawClaim(c context.Context, msg *types.MsgWithdrawClaim) (*types.MsgWithdrawClaimResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	orch, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
// 	validator := k.GetOrchestratorValidator(ctx, orch)
// 	if validator == nil {
// 		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(orch))
// 		if sval == nil {
// 			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
// 		}
// 		validator = sval.GetOperator()
// 	}

// 	// return an error if the validator isn't in the active set
// 	val := k.StakingKeeper.Validator(ctx, validator)
// 	if val == nil || !val.IsBonded() {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
// 	}

// 	any, err := codectypes.NewAnyWithValue(msg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Add the claim to the store
// 	_, err = k.Attest(ctx, msg, any)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(err, "create attestation")
// 	}

// 	// Emit the handle message event
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
// 			// TODO: maybe return something better here? is this the right string representation?
// 			sdk.NewAttribute(types.AttributeKeyAttestationID, string(types.GetAttestationKey(msg.EventNonce, msg.ClaimHash()))),
// 		),
// 	)

// 	return &types.MsgWithdrawClaimResponse{}, nil
// }

// // ERC20Deployed handles MsgERC20Deployed
// func (k msgServer) ERC20DeployedClaim(c context.Context, msg *types.MsgERC20DeployedClaim) (*types.MsgERC20DeployedClaimResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	orch, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
// 	validator := k.GetOrchestratorValidator(ctx, orch)
// 	if validator == nil {
// 		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(orch))
// 		if sval == nil {
// 			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
// 		}
// 		validator = sval.GetOperator()
// 	}

// 	// return an error if the validator isn't in the active set
// 	val := k.StakingKeeper.Validator(ctx, validator)
// 	if val == nil || !val.IsBonded() {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
// 	}

// 	any, err := codectypes.NewAnyWithValue(msg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Add the claim to the store
// 	_, err = k.Attest(ctx, msg, any)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(err, "create attestation")
// 	}

// 	// Emit the handle message event
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
// 			// TODO: maybe return something better here? is this the right string representation?
// 			sdk.NewAttribute(types.AttributeKeyAttestationID, string(types.GetAttestationKey(msg.EventNonce, msg.ClaimHash()))),
// 		),
// 	)

// 	return &types.MsgERC20DeployedClaimResponse{}, nil
// }

// SubmitClaim handles claims for executing a logic call on Ethereum
func (k msgServer) SubmitClaim(c context.Context, msg *types.MsgSubmitClaim) (*types.MsgSubmitClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	orch, _ := sdk.AccAddressFromBech32(msg.Signer)
	validator := k.GetOrchestratorValidator(ctx, orch)
	if validator == nil {
		sval := k.StakingKeeper.Validator(ctx, sdk.ValAddress(orch))
		if sval == nil {
			return nil, sdkerrors.Wrap(types.ErrUnknown, "validator")
		}
		validator = sval.GetOperator()
	}

	// return an error if the validator isn't in the active set
	val := k.StakingKeeper.Validator(ctx, validator)
	if val == nil || !val.IsBonded() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "validator not in acitve set")
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	// Add the claim to the store
	_, err = k.Attest(ctx, msg, any)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create attestation")
	}

	// Emit the handle message event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			// TODO: maybe return something better here? is this the right string representation?
			sdk.NewAttribute(types.AttributeKeyAttestationID, string(types.GetAttestationKey(msg.EventNonce, msg.ClaimHash()))),
		),
	)

	return &types.MsgSubmitClaimResponse{}, nil
}

func (k msgServer) CancelSendToEth(c context.Context, msg *types.MsgCancelSendToEth) (*types.MsgCancelSendToEthResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	err = k.RemoveFromOutgoingPoolAndRefund(ctx, msg.TransactionId, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyOutgoingTXID, fmt.Sprint(msg.TransactionId)),
		),
	)

	return &types.MsgCancelSendToEthResponse{}, nil
}
