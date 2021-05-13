package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "gravity"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey is the module name router key
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

// TODO: go over these key names to fix the find and replace carnage
var (
	// EthAddressByValidatorKey indexes cosmos validator account addresses
	// i.e. cosmos1ahx7f8wyertuus9r20284ej0asrs085case3kn
	EthAddressByValidatorKey = []byte{0x1}

	// ValidatorByEthAddressKey indexes ethereum addresses
	// i.e. 0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B
	ValidatorByEthAddressKey = []byte{0x2}

	// SignerSetTxKey indexes signer set tx by nonce
	SignerSetTxKey = []byte{0x3}

	// SignerSetTxSignatureKey indexes signer set tx signatures by nonce and the validator account address
	// i.e cosmos1ahx7f8wyertuus9r20284ej0asrs085case3kn
	SignerSetTxSignatureKey = []byte{0x4}

	// EthereumEventKey by nonce and validator address
	// i.e. cosmosvaloper1ahx7f8wyertuus9r20284ej0asrs085case3kn
	// A event is named more intuitively than an EthereumEventVoteRecord, it is literally
	// a validator making a event to have seen something happen. Claims are
	// attached to ethereumEventVoteRecords which can be thought of as 'the event' that
	// will eventually be executed.
	EthereumEventKey = []byte{0x5}

	// EthereumEventVoteRecordKey ethereumEventVoteRecord details by nonce and validator address
	// i.e. cosmosvaloper1ahx7f8wyertuus9r20284ej0asrs085case3kn
	// An ethereumEventVoteRecord can be thought of as the 'event to be executed' while
	// the Claims are an individual validator saying that they saw an event
	// occur the EthereumEventVoteRecord is 'the event' that multiple events vote on and
	// eventually executes
	EthereumEventVoteRecordKey = []byte{0x6}

	// OutgoingTXPoolKey indexes the last nonce for the outgoing tx pool
	OutgoingTXPoolKey = []byte{0x7}

	// DenomiatorPrefix indexes token contract addresses from ETH on gravity
	DenomiatorPrefix = []byte{0x8}

	// SecondIndexOutgoingTXFeeKey indexes fee amounts by token contract address
	SecondIndexOutgoingTXFeeKey = []byte{0x9}

	// OutgoingTXBatchKey indexes outgoing tx batches under a nonce and token address
	OutgoingTXBatchKey = []byte{0xa}

	// OutgoingTXBatchBlockKey indexes outgoing tx batches under a block height and token address
	OutgoingTXBatchBlockKey = []byte{0xb}

	// BatchTxSignatureKey indexes batch tx signatures by token contract address
	BatchTxSignatureKey = []byte{0xc}

	// SecondIndexNonceByClaimKey indexes latest nonce for a given event type
	// TOOD: this isn't being used
	SecondIndexNonceByClaimKey = []byte{0xd}

	// LastEventNonceByValidatorKey indexes lateset event nonce by validator
	LastEventNonceByValidatorKey = []byte{0xe}

	// LastObservedEventNonceKey indexes the latest event nonce
	LastObservedEventNonceKey = []byte{0xf}

	// SequenceKeyPrefix indexes different txids
	SequenceKeyPrefix = []byte{0x10}

	// KeyLastTXPoolID indexes the lastTxPoolID
	KeyLastTXPoolID = append(SequenceKeyPrefix, []byte("lastTxPoolId")...)

	// KeyLastOutgoingBatchID indexes the lastBatchID
	KeyLastOutgoingBatchID = append(SequenceKeyPrefix, []byte("lastBatchId")...)

	// KeyOrchestratorAddress indexes the validator keys for an orchestrator
	KeyOrchestratorAddress = []byte{0x11}

	// KeyContractCallTx indexes the outgoing logic calls
	KeyContractCallTx = []byte{0x12}

	// KeyContractCallTxSignature indexes the contract call tx signatures
	KeyContractCallTxSignature = []byte{0x13}

	// DenomToERC20Key prefixes the index of Cosmos originated asset denoms to ERC20s
	DenomToERC20Key = []byte{0x15}

	// ERC20ToDenomKey prefixes the index of Cosmos originated assets ERC20s to denoms
	ERC20ToDenomKey = []byte{0x16}

	// LastSlashedSignerSetTxNonce indexes the latest slashed valset nonce
	LastSlashedSignerSetTxNonce = []byte{0x17}

	// LatestSignerSetTxNonce indexes the latest valset nonce
	LatestSignerSetTxNonce = []byte{0x18}

	// LastSlashedBatchBlock indexes the latest slashed batch block height
	LastSlashedBatchBlock = []byte{0x19}

	// LastUnBondingBlockHeight indexes the last validator unbonding block height
	LastUnBondingBlockHeight = []byte{0xf8}

	// LatestEthereumBlockHeightKey indexes the latest Ethereum block height
	LatestEthereumBlockHeightKey = []byte{0xf9}

	// LastObservedSignerSetTxNonceKey indexes the latest observed valset nonce
	// HERE THERE BE DRAGONS, do not use this value as an up to date validator set
	// on Ethereum it will always lag significantly and may be totally wrong at some
	// times.
	LastObservedSignerSetTxKey = []byte{0xfa}

	// PastEthSignatureCheckpointKey indexes eth signature checkpoints that have existed
	PastEthSignatureCheckpointKey = []byte{0x1b}
)

// GetOrchestratorAddressKey returns the following key format
// prefix
// [0xe8][cosmos1ahx7f8wyertuus9r20284ej0asrs085case3kn]
func GetOrchestratorAddressKey(orc sdk.AccAddress) []byte {
	return append(KeyOrchestratorAddress, orc.Bytes()...)
}

// GetEthAddressByValidatorKey returns the following key format
// prefix              cosmos-validator
// [0x0][cosmosvaloper1ahx7f8wyertuus9r20284ej0asrs085case3kn]
func GetEthAddressByValidatorKey(validator sdk.ValAddress) []byte {
	return append(EthAddressByValidatorKey, validator.Bytes()...)
}

// GetValidatorByEthAddressKey returns the following key format
// prefix              cosmos-validator
// [0xf9][0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B]
func GetValidatorByEthAddressKey(ethAddress string) []byte {
	return append(ValidatorByEthAddressKey, []byte(ethAddress)...)
}

// GetSignerSetTxKey returns the following key format
// prefix    nonce
// [0x0][0 0 0 0 0 0 0 1]
func GetSignerSetTxKey(nonce uint64) []byte {
	return append(SignerSetTxKey, UInt64Bytes(nonce)...)
}

// GetSignerSetTxSignatureKey returns the following key format
// prefix   nonce                    validator-address
// [0x0][0 0 0 0 0 0 0 1][cosmos1ahx7f8wyertuus9r20284ej0asrs085case3kn]
// MARK finish-batches: this is where the key is created in the old (presumed working) code
func GetSignerSetTxSignatureKey(nonce uint64, validator sdk.AccAddress) []byte {
	return append(SignerSetTxSignatureKey, append(UInt64Bytes(nonce), validator.Bytes()...)...)
}

// GetClaimKey returns the following key format
// prefix type               cosmos-validator-address                       nonce                             ethereumEventVoteRecord-details-hash
// [0x0][0 0 0 1][cosmosvaloper1ahx7f8wyertuus9r20284ej0asrs085case3kn][0 0 0 0 0 0 0 1][fd1af8cec6c67fcf156f1b61fdf91ebc04d05484d007436e75342fc05bbff35a]
// The Claim hash identifies a unique event, for example it would have a event nonce, a sender and a receiver. Or an event nonce and a batch nonce. But
// the Claim is stored indexed with the eventer key to make sure that it is unique.
// TODO: This is dead
func GetClaimKey(details EthereumEvent) []byte {
	var detailsHash []byte
	if details != nil {
		detailsHash = details.EventHash()
	} else {
		panic("No event without details!")
	}
	eventTypeLen := len([]byte{byte(details.GetType())})
	nonceBz := UInt64Bytes(details.GetEventNonce())
	key := make([]byte, len(EthereumEventKey)+eventTypeLen+sdk.AddrLen+len(nonceBz)+len(detailsHash))
	copy(key[0:], EthereumEventKey)
	copy(key[len(EthereumEventKey):], []byte{byte(details.GetType())})
	// TODO this is the delegate address, should be stored by the valaddress
	copy(key[len(EthereumEventKey)+eventTypeLen:], details.GetValidator())
	copy(key[len(EthereumEventKey)+eventTypeLen+sdk.AddrLen:], nonceBz)
	copy(key[len(EthereumEventKey)+eventTypeLen+sdk.AddrLen+len(nonceBz):], detailsHash)
	return key
}

// GetEthereumEventVoteRecordKey returns the following key format
// prefix     nonce                             event-details-hash
// [0x5][0 0 0 0 0 0 0 1][fd1af8cec6c67fcf156f1b61fdf91ebc04d05484d007436e75342fc05bbff35a]
// An ethereumEventVoteRecord is an event multiple people are voting on, this function needs the event
// details because each EthereumEventVoteRecord is aggregating all events of a specific event, lets say
// validator X and validator y where making different events about the same event nonce
// Note that the event hash does NOT include the eventer address and only identifies an event
func GetEthereumEventVoteRecordKey(eventNonce uint64, eventHash []byte) []byte {
	key := make([]byte, len(EthereumEventVoteRecordKey)+len(UInt64Bytes(0))+len(eventHash))
	copy(key[0:], EthereumEventVoteRecordKey)
	copy(key[len(EthereumEventVoteRecordKey):], UInt64Bytes(eventNonce))
	copy(key[len(EthereumEventVoteRecordKey)+len(UInt64Bytes(0)):], eventHash)
	return key
}

// GetEthereumEventVoteRecordKeyWithHash returns the following key format
// prefix     nonce                             event-details-hash
// [0x5][0 0 0 0 0 0 0 1][fd1af8cec6c67fcf156f1b61fdf91ebc04d05484d007436e75342fc05bbff35a]
// An ethereumEventVoteRecord is an event multiple people are voting on, this function needs the event
// details because each EthereumEventVoteRecord is aggregating all events of a specific event, lets say
// validator X and validator y where making different events about the same event nonce
// Note that the event hash does NOT include the eventer address and only identifies an event
func GetEthereumEventVoteRecordKeyWithHash(eventNonce uint64, eventHash []byte) []byte {
	key := make([]byte, len(EthereumEventVoteRecordKey)+len(UInt64Bytes(0))+len(eventHash))
	copy(key[0:], EthereumEventVoteRecordKey)
	copy(key[len(EthereumEventVoteRecordKey):], UInt64Bytes(eventNonce))
	copy(key[len(EthereumEventVoteRecordKey)+len(UInt64Bytes(0)):], eventHash)
	return key
}

// GetOutgoingTxPoolKey returns the following key format
// prefix     id
// [0x6][0 0 0 0 0 0 0 1]
func GetOutgoingTxPoolKey(id uint64) []byte {
	return append(OutgoingTXPoolKey, sdk.Uint64ToBigEndian(id)...)
}

// GetBatchTxKey returns the following key format
// prefix     nonce                     eth-contract-address
// [0xa][0 0 0 0 0 0 0 1][0xc783df8a850f42e7F7e57013759C285caa701eB6]
func GetBatchTxKey(tokenContract string, nonce uint64) []byte {
	return append(append(OutgoingTXBatchKey, []byte(tokenContract)...), UInt64Bytes(nonce)...)
}

// GetBatchTxBlockKey returns the following key format
// prefix     blockheight
// [0xb][0 0 0 0 2 1 4 3]
func GetBatchTxBlockKey(block uint64) []byte {
	return append(OutgoingTXBatchBlockKey, UInt64Bytes(block)...)
}

// GetBatchTxSignatureKey returns the following key format
// prefix           eth-contract-address                BatchNonce                       Validator-address
// [0xe1][0xc783df8a850f42e7F7e57013759C285caa701eB6][0 0 0 0 0 0 0 1][cosmosvaloper1ahx7f8wyertuus9r20284ej0asrs085case3kn]
// TODO this should be a sdk.ValAddress
func GetBatchTxSignatureKey(tokenContract string, batchNonce uint64, validator sdk.AccAddress) []byte {
	a := append(UInt64Bytes(batchNonce), validator.Bytes()...)
	b := append([]byte(tokenContract), a...)
	c := append(BatchTxSignatureKey, b...)
	return c
}

// GetFeeSecondIndexKey returns the following key format
// prefix            eth-contract-address            fee_amount
// [0x9][0xc783df8a850f42e7F7e57013759C285caa701eB6][1000000000]
func GetFeeSecondIndexKey(fee ERC20Token) []byte {
	r := make([]byte, 1+ETHContractAddressLen+32)
	// sdkInts have a size limit of 255 bits or 32 bytes
	// therefore this will never panic and is always safe
	amount := make([]byte, 32)
	amount = fee.Amount.BigInt().FillBytes(amount)
	// TODO this won't ever work fix it
	copy(r[0:], SecondIndexOutgoingTXFeeKey)
	copy(r[len(SecondIndexOutgoingTXFeeKey):], []byte(fee.Contract))
	copy(r[len(SecondIndexOutgoingTXFeeKey)+len(fee.Contract):], amount)
	return r
}

// GetLastEventNonceByValidatorKey indexes lateset event nonce by validator
// GetLastEventNonceByValidatorKey returns the following key format
// prefix              cosmos-validator
// [0x0][cosmos1ahx7f8wyertuus9r20284ej0asrs085case3kn]
func GetLastEventNonceByValidatorKey(validator sdk.ValAddress) []byte {
	return append(LastEventNonceByValidatorKey, validator.Bytes()...)
}

func GetDenomToERC20Key(denom string) []byte {
	return append(DenomToERC20Key, []byte(denom)...)
}

func GetERC20ToDenomKey(erc20 string) []byte {
	return append(ERC20ToDenomKey, []byte(erc20)...)
}

func GetContractCallTxKey(invalidationId []byte, invalidationNonce uint64) []byte {
	a := append(KeyContractCallTx, invalidationId...)
	return append(a, UInt64Bytes(invalidationNonce)...)
}

func GetLogicConfirmKey(invalidationId []byte, invalidationNonce uint64, validator sdk.AccAddress) []byte {
	interm := append(KeyContractCallTxSignature, invalidationId...)
	interm = append(interm, UInt64Bytes(invalidationNonce)...)
	return append(interm, validator.Bytes()...)
}

// GetPastEthSignatureCheckpointKey returns the following key format
// prefix    checkpoint
// [0x0][ checkpoint bytes ]
func GetPastEthSignatureCheckpointKey(checkpoint []byte) []byte {
	return append(PastEthSignatureCheckpointKey, checkpoint...)
}
