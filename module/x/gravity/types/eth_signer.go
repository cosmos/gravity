package types

import (
	"fmt"
	"math/big"
	"sort"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ValidateBasic performs stateless checks on validity
func (s EthSigner) ValidateBasic() error {
	if s.Power == 0 {
		return fmt.Errorf("consensus power cannot be 0")
	}
	if err := ValidateEthAddress(s.EthereumAddress); err != nil {
		return fmt.Errorf("invalid signer ethereum address: %w", err)
	}
	return nil
}

// EthSigners is the sorted set of validator data for Ethereum bridge MultiSig set
type EthSigners []EthSigner

// Sort sorts the validators by power
func (s EthSigners) Sort() {
	sort.Slice(s, func(i, j int) bool {
		if s[i].Power == s[j].Power {
			// Secondary sort on eth address in case powers are equal
			return s[i].EthereumAddress > s[j].EthereumAddress
		}
		return s[i].Power > s[j].Power
	})
}

// TotalPower returns the total power in the bridge validator set
func (s EthSigners) TotalPower() int64 {
	totalPower := int64(0)
	for _, signer := range s {
		totalPower += signer.Power
	}
	return totalPower
}

// ValidateBasic performs stateless checks
func (s EthSigners) ValidateBasic() error {
	seenSigners := make(map[string]bool, 0)
	for _, signer := range s {
		if seenSigners[signer.EthereumAddress] {
			return fmt.Errorf("duplicate entry for signer %s", signer.EthereumAddress)
		}
		if err := signer.ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "signer %s validation failed", signer.EthereumAddress)
		}
		seenSigners[signer.EthereumAddress] = true
	}

	return nil
}

// NewSignerSet returns a new ethereum signer set based on the staking bonded
// validator set
func NewSignerSet(height uint64, signers ...EthSigner) EthSignerSet {
	ethSigners := EthSigners(signers)
	ethSigners.Sort()

	return EthSignerSet{
		Signers: ethSigners,
		Height:  height,
	}
}

// GetCheckpoint returns the checkpoint
func (ss EthSignerSet) GetCheckpoint(bridgeID []byte) ([]byte, error) {
	contractABI, err := abi.JSON(strings.NewReader(SignerSetCheckpointABIJSON))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "bad ABI definition in code")
	}

	checkpointBytes := []uint8("checkpoint")
	var checkpoint [32]uint8
	copy(checkpoint[:], checkpointBytes[:])

	memberAddresses := make([]common.Address, len(ss.Signers))
	convertedPowers := make([]*big.Int, len(ss.Signers))
	for i, m := range ss.Signers {
		memberAddresses[i] = common.HexToAddress(m.EthereumAddress)
		convertedPowers[i] = big.NewInt(int64(m.Power))
	}
	// the word 'checkpoint' needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	bytes, err := contractABI.Pack(
		"checkpoint",
		bridgeID,
		checkpoint,
		big.NewInt(int64(ss.Height)),
		memberAddresses,
		convertedPowers,
	)

	if err != nil {
		// this should never happen outside of test since any case that could crash on encoding
		// should be filtered above.
		return nil, sdkerrors.Wrap(err, "packing checkpoint")
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you where to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	hash := crypto.Keccak256Hash(bytes[4:])
	return hash.Bytes(), err
}
