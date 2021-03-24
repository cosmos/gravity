package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Confirm interface {
	GetType() ConfirmType
	GetOrchestratorAddress() string
	GetNonce() uint64
	GetSignature() string
}

var (
	_ Confirm = &ConfirmBatch{}
	_ Confirm = &ConfirmLogicCall{}
	_ Confirm = &ValsetConfirm{}
)

// Type should return the action
func (msg ConfirmBatch) GetType() ConfirmType { return ConfirmType_CONFIRM_TYPE_BATCH }

// ValidateBasic performs stateless checks
func (msg ConfirmBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.OrchestratorAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OrchestratorAddress)
	}
	if err := ValidateEthAddress(msg.EthSigner); err != nil {
		return sdkerrors.Wrap(err, "eth signer")
	}
	if err := ValidateEthAddress(msg.TokenContract); err != nil {
		return sdkerrors.Wrap(err, "token contract")
	}
	_, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Could not decode hex string %s", msg.Signature)
	}
	return nil
}

// Type should return the action
func (msg ConfirmLogicCall) GetType() ConfirmType { return ConfirmType_CONFIRM_TYPE_LOGIC }

// ValidateBasic performs stateless checks
func (msg ConfirmLogicCall) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.OrchestratorAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OrchestratorAddress)
	}
	if err := ValidateEthAddress(msg.EthSigner); err != nil {
		return sdkerrors.Wrap(err, "eth signer")
	}
	_, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Could not decode hex string %s", msg.Signature)
	}
	_, err = hex.DecodeString(msg.InvalidationId)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Could not decode hex string %s", msg.InvalidationId)
	}
	return nil
}

func (msg ConfirmLogicCall) GetNonce() uint64 {
	return 0
}

// NewMsgValsetConfirm returns a new ValsetConfirm
func NewMsgValsetConfirm(nonce uint64, ethAddress string, validator sdk.AccAddress, signature string) *ValsetConfirm {
	return &ValsetConfirm{
		Nonce:               nonce,
		OrchestratorAddress: validator.String(),
		EthAddress:          ethAddress,
		Signature:           signature,
	}
}

// Type should return the action
func (msg *ValsetConfirm) GetType() ConfirmType { return ConfirmType_CONFIRM_TYPE_VALSET }

// ValidateBasic performs stateless checks
func (msg *ValsetConfirm) ValidateBasic() (err error) {
	if _, err = sdk.AccAddressFromBech32(msg.OrchestratorAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OrchestratorAddress)
	}
	if err := ValidateEthAddress(msg.EthAddress); err != nil {
		return sdkerrors.Wrap(err, "ethereum address")
	}
	return nil
}
