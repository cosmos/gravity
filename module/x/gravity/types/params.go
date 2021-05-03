package types

import (
	"errors"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
)

// DefaultParamspace defines the default gravity module parameter subspace
const DefaultParamspace = ModuleName

var (
	// ParamsStoreKeyBridgeContractAddress stores the contract address
	ParamsStoreKeyBridgeContractAddress = []byte("BridgeContractAddress")

	// ParamsStoreKeyBridgeContractChainID stores the bridge chain id
	ParamsStoreKeyBridgeContractChainID = []byte("BridgeChainID")

	// ParamsStoreKeySignerSetWindow stores the signed blocks window
	ParamsStoreKeySignerSetWindow = []byte("SignerSetWindow")

	// ParamsStoreKeyBatchTxWindow stores the signed blocks window
	ParamsStoreKeyBatchTxWindow = []byte("BatchTxWindow")

	// ParamsStoreKeyEventWindow stores the signed blocks window
	ParamsStoreKeyEventWindow = []byte("EventWindow")

	// ParamsStoreKeyTargetBatchTimeout stores the signed blocks window
	ParamsStoreKeyTargetBatchTimeout = []byte("TargetBatchTimeout")

	// ParamsStoreKeyAverageBlockTime stores the signed blocks window
	ParamsStoreKeyAverageBlockTime = []byte("AverageBlockTime")

	// ParamsStoreKeyBatchSize stores the batch size
	ParamsStoreKeyBatchSize = []byte("BatchSize")

	// ParamsStoreKeyAverageEthereumBlockTime stores the signed blocks window
	ParamsStoreKeyAverageEthereumBlockTime = []byte("AverageEthereumBlockTime")

	// ParamsStoreSlashFractionSignerSet stores the slash fraction signer set
	ParamsStoreSlashFractionSignerSet = []byte("SlashFractionSignerSet")

	// ParamsStoreSlashFractionBatch stores the slash fraction batch
	ParamsStoreSlashFractionBatch = []byte("SlashFractionBatch")

	// ParamsStoreSlashFractionLogicCall stores the slash fraction logic call
	ParamsStoreSlashFractionLogicCall = []byte("SlashFractionLogicCall")

	// ParamsStoreSlashFractionEvent stores the slash fraction Claim
	ParamsStoreSlashFractionEvent = []byte("SlashFractionEvent")

	// ParamsStoreSlashFractionConflictingEvent stores the slash fraction ConflictingClaim
	ParamsStoreSlashFractionConflictingEvent = []byte("SlashFractionConflictingEvent")

	// ParamStoreUnbondingWindow stores unbond slashing valset window
	ParamStoreUnbondingWindow = []byte("UnbondingWindow")

	// ParamStoreMaxSignerSetPowerDiff stores the power diff threshold value for
	// the ethereum signer sets
	ParamStoreMaxSignerSetPowerDiff = []byte("MaxSignerSetPowerDiff")

	// ParamsStoreKeyAttestationVotesPowerThreshold stores the attestation voting power threshold
	ParamsStoreKeyAttestationVotesPowerThreshold = []byte("AttestationVotesPowerThreshold")

	// Ensure that params implements the proper interface
	_ paramtypes.ParamSet = &Params{}
)

// DefaultParams returns a copy of the default params
func DefaultParams() Params {
	return Params{
		BridgeContractAddress:          common.Address{}.String(),
		BridgeChainID:                  1, // Ethereum Mainnet
		SignerSetWindow:                10000,
		BatchTxWindow:                  10000,
		EventWindow:                    10000,
		TargetBatchTimeout:             43200000,
		AverageBlockTime:               5000,
		AverageEthereumBlockTime:       15000,
		BatchSize:                      100,
		SlashFractionSignerSet:         sdk.NewDecWithPrec(1, 3), // 0.1 %
		SlashFractionBatch:             sdk.NewDecWithPrec(1, 3), // 0.1 %
		SlashFractionLogicCall:         sdk.NewDecWithPrec(1, 3), // 0.1 %
		SlashFractionEvent:             sdk.NewDecWithPrec(1, 3), // 0.1 %
		SlashFractionConflictingEvent:  sdk.NewDecWithPrec(1, 3), // 0.1 %
		UnbondingWindow:                10000,
		MaxSignerSetPowerDiff:          sdk.NewDecWithPrec(5, 2),   // 5 %
		AttestationVotesPowerThreshold: sdk.NewDecWithPrec(666, 3), // 66.6 %
	}
}

// ValidateBasic checks that the parameters have valid values.
func (p Params) ValidateBasic() error {
	// TODO: reuse validation functions
	if err := validateBridgeContractAddress(p.BridgeContractAddress); err != nil {
		return sdkerrors.Wrap(err, "bridge contract address")
	}
	if err := validateBridgeChainID(p.BridgeChainID); err != nil {
		return sdkerrors.Wrap(err, "bridge chain id")
	}
	if err := validateTargetBatchTimeout(p.TargetBatchTimeout); err != nil {
		return sdkerrors.Wrap(err, "target batch timeout")
	}
	if err := validateAverageBlockTime(p.AverageBlockTime); err != nil {
		return sdkerrors.Wrap(err, "avg block time")
	}
	if err := validateAverageEthereumBlockTime(p.AverageEthereumBlockTime); err != nil {
		return sdkerrors.Wrap(err, "avg ethereum block time")
	}
	if err := validateBatchSize(p.BatchSize); err != nil {
		return sdkerrors.Wrap(err, "batch size")
	}
	if err := validateSignerSetWindow(p.SignerSetWindow); err != nil {
		return sdkerrors.Wrap(err, "signer set rolling window")
	}
	if err := validateBatchTxWindow(p.BatchTxWindow); err != nil {
		return sdkerrors.Wrap(err, "batch tx rolling window")
	}
	if err := validateEventWindow(p.EventWindow); err != nil {
		return sdkerrors.Wrap(err, "event rolling window")
	}
	if err := validateSlashFractionSignerSet(p.SlashFractionSignerSet); err != nil {
		return sdkerrors.Wrap(err, "slash fraction valset")
	}
	if err := validateSlashFractionBatch(p.SlashFractionBatch); err != nil {
		return sdkerrors.Wrap(err, "slash fraction batch")
	}
	// FIXME: update func
	if err := validateSlashFractionBatch(p.SlashFractionLogicCall); err != nil {
		return sdkerrors.Wrap(err, "slash fraction logic call")
	}
	if err := validateSlashFractionEvent(p.SlashFractionEvent); err != nil {
		return sdkerrors.Wrap(err, "slash fraction event")
	}
	if err := validateSlashFractionConflictingEvent(p.SlashFractionConflictingEvent); err != nil {
		return sdkerrors.Wrap(err, "slash fraction conflicting event")
	}
	if err := validateUnbondingWindow(p.UnbondingWindow); err != nil {
		return sdkerrors.Wrap(err, "unbond slashing validator window")
	}
	// FIXME: update func
	if err := validateSlashFractionConflictingEvent(p.AttestationVotesPowerThreshold); err != nil {
		return sdkerrors.Wrap(err, "attestation voting power threshold")
	}
	return nil
}

// ParamKeyTable for gravity module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of gravity module's parameters.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamsStoreKeyBridgeContractAddress, &p.BridgeContractAddress, validateBridgeContractAddress),
		paramtypes.NewParamSetPair(ParamsStoreKeyBridgeContractChainID, &p.BridgeChainID, validateBridgeChainID),
		paramtypes.NewParamSetPair(ParamsStoreKeySignerSetWindow, &p.SignerSetWindow, validateSignerSetWindow),
		paramtypes.NewParamSetPair(ParamsStoreKeyBatchTxWindow, &p.BatchTxWindow, validateBatchTxWindow),
		paramtypes.NewParamSetPair(ParamsStoreKeyEventWindow, &p.EventWindow, validateEventWindow),
		paramtypes.NewParamSetPair(ParamsStoreKeyAverageBlockTime, &p.AverageBlockTime, validateAverageBlockTime),
		paramtypes.NewParamSetPair(ParamsStoreKeyTargetBatchTimeout, &p.TargetBatchTimeout, validateTargetBatchTimeout),
		paramtypes.NewParamSetPair(ParamsStoreKeyAverageEthereumBlockTime, &p.AverageEthereumBlockTime, validateAverageEthereumBlockTime),
		paramtypes.NewParamSetPair(ParamsStoreKeyBatchSize, &p.BatchSize, validateBatchSize),
		paramtypes.NewParamSetPair(ParamsStoreSlashFractionSignerSet, &p.SlashFractionSignerSet, validateSlashFractionSignerSet),
		paramtypes.NewParamSetPair(ParamsStoreSlashFractionBatch, &p.SlashFractionBatch, validateSlashFractionBatch),
		paramtypes.NewParamSetPair(ParamsStoreSlashFractionBatch, &p.SlashFractionLogicCall, validateSlashFractionBatch),
		paramtypes.NewParamSetPair(ParamsStoreSlashFractionEvent, &p.SlashFractionEvent, validateSlashFractionEvent),
		paramtypes.NewParamSetPair(ParamsStoreSlashFractionConflictingEvent, &p.SlashFractionConflictingEvent, validateSlashFractionConflictingEvent),
		paramtypes.NewParamSetPair(ParamStoreUnbondingWindow, &p.UnbondingWindow, validateUnbondingWindow),
		paramtypes.NewParamSetPair(ParamStoreMaxSignerSetPowerDiff, &p.MaxSignerSetPowerDiff, validateSlashFractionSignerSet),
		paramtypes.NewParamSetPair(ParamsStoreKeyAttestationVotesPowerThreshold, &p.AttestationVotesPowerThreshold, validateSlashFractionSignerSet),
	}
}

var ErrWrongType = errors.New("invalid type")
var ErrTooLarge = errors.New("value above bound")
var ErrTooSmall = errors.New("value below bound")

func validateBoundedUInt64(i interface{}, lower uint64, upper uint64) error {
	u, ok := i.(uint64)
	if !ok {
		return sdkerrors.Wrapf(ErrWrongType, fmt.Sprintf("invalid parameter type: %T", i))
	}

	if u > upper {
		return sdkerrors.Wrapf(ErrTooLarge, fmt.Sprintf("parameter value %d larger than bound %d", u, upper))
	}

	if u < lower {
		return sdkerrors.Wrapf(ErrTooSmall, fmt.Sprintf("parameter value %d smaller than bound %d", u, lower))
	}

	return nil
}

func validateBoundedDec(i interface{}, lower sdk.Dec, upper sdk.Dec) error {
	d, ok := i.(sdk.Dec)
	if !ok {
		return sdkerrors.Wrapf(ErrWrongType, fmt.Sprintf("invalid parameter type: %T", i))
	}

	if d.GT(upper) {
		return sdkerrors.Wrapf(ErrTooLarge, fmt.Sprintf("parameter value %s larger than bound %s", d, upper))
	}

	if d.LT(lower) {
		return sdkerrors.Wrapf(ErrTooSmall, fmt.Sprintf("parameter value %s smaller than bound %s", d, lower))
	}

	return nil
}

func validateBridgeChainID(i interface{}) error {
	chainID, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if chainID == 0 {
		return fmt.Errorf("invalid chain ID %d", chainID)
	}
	return nil
}

func validateTargetBatchTimeout(i interface{}) error {
	err := validateBoundedUInt64(i, 60000, math.MaxUint64)
	if err != nil {
		if err == ErrTooSmall {
			return sdkerrors.Wrapf(err, "invalid target batch timeout, less than 60 seconds is too short")
		}
	}
	return err
}

func validateAverageBlockTime(i interface{}) error {
	err := validateBoundedUInt64(i, 100, math.MaxUint64)
	if err != nil {
		if err == ErrTooSmall {
			return sdkerrors.Wrapf(err, "invalid average Cosmos block time, too short for latency limitations")
		}
	}
	return err
}

func validateAverageEthereumBlockTime(i interface{}) error {
	err := validateBoundedUInt64(i, 100, math.MaxUint64)
	if err != nil {
		if err == ErrTooSmall {
			return sdkerrors.Wrapf(err, "invalid average Ethereum block time, too short for latency limitations")
		}
	}
	return err
}

func validateBatchSize(i interface{}) error {
	err := validateBoundedUInt64(i, 1, math.MaxUint64)
	if err != nil {
		if err == ErrTooSmall {
			return sdkerrors.Wrapf(err, "batch tx size cannot be 0")
		}
	}
	return err
}

func validateBridgeContractAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return ValidateEthAddress(v)
}

func validateSignerSetWindow(i interface{}) error {
	return validateBoundedUInt64(i, 0, math.MaxUint64)
}

func validateUnbondingWindow(i interface{}) error {
	return validateBoundedUInt64(i, 0, math.MaxUint64)
}

func validateSlashFractionSignerSet(i interface{}) error {
	return validateBoundedDec(i, sdk.ZeroDec(), sdk.OneDec())
}

func validateBatchTxWindow(i interface{}) error {
	return validateBoundedUInt64(i, 0, math.MaxUint64)
}

func validateEventWindow(i interface{}) error {
	return validateBoundedUInt64(i, 0, math.MaxUint64)
}

func validateSlashFractionBatch(i interface{}) error {
	return validateBoundedDec(i, sdk.ZeroDec(), sdk.OneDec())
}

func validateSlashFractionEvent(i interface{}) error {
	return validateBoundedDec(i, sdk.ZeroDec(), sdk.OneDec())
}

func validateSlashFractionConflictingEvent(i interface{}) error {
	return validateBoundedDec(i, sdk.ZeroDec(), sdk.OneDec())
}
