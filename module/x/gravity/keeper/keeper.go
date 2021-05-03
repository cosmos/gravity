package keeper

import (
	"crypto/sha256"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/gravity-bridge/module/x/gravity/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace

	cdc            codec.BinaryMarshaler
	bankKeeper     types.BankKeeper
	slashingKeeper types.SlashingKeeper
	stakingKeeper  types.StakingKeeper

	attestationHandler AttestationHandler
}

// NewKeeper returns a new instance of the gravity keeper
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	stakingKeeper types.StakingKeeper, bankKeeper types.BankKeeper, slashingKeeper types.SlashingKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		paramSpace:     paramSpace,
		storeKey:       storeKey,
		bankKeeper:     bankKeeper,
		slashingKeeper: slashingKeeper,
		stakingKeeper:  stakingKeeper,
	}
}

// SetAttestationHandler sets an attestation handler for the bridge module. This function panics if
// the attestation handler is already set
func (k *Keeper) SetAttestationHandler(handler AttestationHandler) {
	if k.attestationHandler != nil {
		panic("attestation handler already set")
	}

	if handler == nil {
		panic("attestation handler provided cannot be nil")
	}

	k.attestationHandler = handler
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetBridgeID returns the BridgeID.
func (k Keeper) GetBridgeID(ctx sdk.Context) tmbytes.HexBytes {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.BridgeIDKey)
}

// SetBridgeID sets the BridgeID value to store
func (k Keeper) SetBridgeID(ctx sdk.Context, ID tmbytes.HexBytes) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BridgeIDKey, ID)
}

// GetEthAddress returns the eth address for a given gravity validatorAddr
func (k Keeper) GetEthAddress(ctx sdk.Context, validatorAddr sdk.ValAddress) common.Address {
	// TODO: use prefix store
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetEthAddressKey(validatorAddr))
	if len(bz) == 0 {
		// return zero address
		return common.Address{}
	}

	return common.BytesToAddress(bz)
}

// SetEthAddress sets the ethereum address for a given validator
func (k Keeper) SetEthAddress(ctx sdk.Context, validatorAddr sdk.ValAddress, ethereumAddr common.Address) {
	// TODO: use prefix store
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetEthAddressKey(validatorAddr), ethereumAddr.Bytes())
}

// GetOrchestratorValidator returns the validator key associated with an orchestrator key
func (k Keeper) GetOrchestratorValidator(ctx sdk.Context, orch sdk.AccAddress) sdk.ValAddress {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetOrchestratorAddressKey(orch))
	if len(bz) == 0 {
		return nil
	}

	return sdk.ValAddress(bz)
}

// SetOrchestratorValidator sets the Orchestrator key for a given validator
func (k Keeper) SetOrchestratorValidator(ctx sdk.Context, val sdk.ValAddress, orch sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetOrchestratorAddressKey(orch), val)
}

// GetEthereumInfo returns the ethereum block height and timestamp of the last
// observed attestation.
func (k Keeper) GetEthereumInfo(ctx sdk.Context) (types.EthereumInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastObservedEthereumBlockHeightKey)
	if len(bz) == 0 {
		return types.EthereumInfo{}, false
	}

	var info types.EthereumInfo
	k.cdc.UnmarshalBinaryBare(bz, &info)
	return info, false
}

// SetEthereumInfo sets an observed ethereum block height and timestamp to the store.
func (k Keeper) SetEthereumInfo(ctx sdk.Context, info types.EthereumInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastObservedEthereumBlockHeightKey, k.cdc.MustMarshalBinaryBare(&info))
}

// GetLastObservedEventNonce returns the latest observed event nonce
func (k Keeper) GetLastObservedEventNonce(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.LastObservedEventNonceKey)
	if len(bytes) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bytes)
}

// SetLastObservedEventNonce sets the latest observed event nonce
func (k Keeper) SetLastObservedEventNonce(ctx sdk.Context, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastObservedEventNonceKey, sdk.Uint64ToBigEndian(nonce))
}

func (k Keeper) GetTransferTx(ctx sdk.Context, id tmbytes.HexBytes) (types.TransferTx, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTransferTxKey(id))
	if len(bz) == 0 {
		return types.TransferTx{}, false
	}

	var tx types.TransferTx
	k.cdc.UnmarshalBinaryBare(bz, &tx)
	return tx, true
}

func (k Keeper) SetTransferTx(ctx sdk.Context, tx types.TransferTx) tmbytes.HexBytes {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TransferTxKey)
	bz := k.cdc.MustMarshalBinaryBare(&tx)
	hash := sha256.Sum256(bz)
	txID := tmbytes.HexBytes(hash[:])
	store.Set(types.GetTransferTxKey(txID), bz)

	return txID
}

func (k Keeper) DeleteTransferTx(ctx sdk.Context, txID tmbytes.HexBytes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TransferTxKey)
	store.Delete(types.GetTransferTxKey(txID))
}

// TODO: update keys and proto
// func (k Keeper) GetTransferBatchMapping(ctx sdk.Context, txID tmbytes.HexBytes) (types.TransferTx, bool) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := store.Get(types.GetTransferTxKey(txID))
// 	if len(bz) == 0 {
// 		return types.TransferTx{}, false
// 	}

// 	var tx types.TransferTx
// 	k.cdc.UnmarshalBinaryBare(bz, &tx)
// 	return tx, true
// }

// func (k Keeper) SetTransferBatchMapping(ctx sdk.Context,txID tmbytes.HexBytes, tokenContract common.Address, batchID tmbytes.HexBytes) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TransferTxKey)
// 	// bz := k.cdc.MustMarshalBinaryBare(&tx)
// 	store.Set(types.GetTransferTxKey(txID), batchID)
// }

// func (k Keeper) DeleteTransferBatchMapping(ctx sdk.Context, txID tmbytes.HexBytes) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TransferTxKey)
// 	store.Delete(types.GetTransferTxKey(txID))
// }

// IterateTransferTxs
func (k Keeper) IterateTransferTxs(ctx sdk.Context, cb func(txID tmbytes.HexBytes, tx types.TransferTx) (stop bool)) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.TransferTxKey)

	iterator := prefixStore.ReverseIterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var tx types.TransferTx
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &tx)
		txID := tmbytes.HexBytes(iterator.Key()[:1]) // TODO: check correctness
		if cb(txID, tx) {
			break // stop iteration
		}
	}
}

// GetTransferTxs returns all the outgoing transactions from the pool in desc order.
// TODO: create struct with ID and transferTx
func (k Keeper) GetTransferTxs(ctx sdk.Context) []types.TransferTx {
	txs := make([]types.TransferTx, 0)
	k.IterateTransferTxs(ctx, func(id tmbytes.HexBytes, tx types.TransferTx) bool {
		txs = append(txs, tx)
		return false
	})

	return txs
}

func (k Keeper) GetEthereumEvent(ctx sdk.Context, eventID tmbytes.HexBytes) (types.EthereumEvent, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EventKeyPrefix)
	bz := store.Get(eventID)
	if len(bz) == 0 {
		return nil, false
	}

	var event types.EthereumEvent
	if err := k.cdc.UnmarshalInterface(bz, event); err != nil {
		panic(err)
	}

	return event, true
}

func (k Keeper) SetEthereumEvent(ctx sdk.Context, eventID tmbytes.HexBytes, event types.EthereumEvent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EventKeyPrefix)
	bz, err := k.cdc.MarshalInterface(event)
	if err != nil {
		panic(err)
	}

	store.Set(eventID, bz)
}

func (k Keeper) IterateEthereumEvents(ctx sdk.Context, cb func(eventID tmbytes.HexBytes, event types.EthereumEvent) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EventKeyPrefix)
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		eventID := tmbytes.HexBytes(iterator.Key()[:1])
		var event types.EthereumEvent
		if err := k.cdc.UnmarshalInterface(iterator.Value(), event); err != nil {
			panic(err)
		}

		if cb(eventID, event) {
			break // stop
		}
	}
}

func (k Keeper) GetConfirm(ctx sdk.Context, confirmID tmbytes.HexBytes) (types.Confirm, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ConfirmKey)
	bz := store.Get(confirmID)
	if len(bz) == 0 {
		return nil, false
	}

	var confirm types.Confirm
	if err := k.cdc.UnmarshalInterface(bz, confirm); err != nil {
		panic(err)
	}

	return confirm, true
}

func (k Keeper) SetConfirm(ctx sdk.Context, confirmID tmbytes.HexBytes, confirm types.Confirm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ConfirmKey)
	bz, err := k.cdc.MarshalInterface(confirm)
	if err != nil {
		panic(err)
	}

	store.Set(confirmID, bz)
}

func (k Keeper) IterateConfirmations(ctx sdk.Context, cb func(confirmID tmbytes.HexBytes, confirm types.Confirm) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ConfirmKey)
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		confirmID := tmbytes.HexBytes(iterator.Key()[:1])
		var confirm types.Confirm
		if err := k.cdc.UnmarshalInterface(iterator.Value(), confirm); err != nil {
			panic(err)
		}

		if cb(confirmID, confirm) {
			break // stop
		}
	}
}

func (k Keeper) IterateValidatorsByPower(ctx sdk.Context, cb func(validator stakingtypes.Validator) bool) {
	iterator := k.stakingKeeper.ValidatorsPowerStoreIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validatorAddr := sdk.ValAddress(iterator.Value())
		validator, found := k.stakingKeeper.GetValidator(ctx, validatorAddr)
		if !found {
			continue
		}

		if cb(validator) {
			break // stop
		}
	}
}
