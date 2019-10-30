package txs

import (
	"context"
	"fmt"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/cosmos/peggy/cmd/ebrelayer/events"
	cosmosBridge "github.com/cosmos/peggy/cmd/ebrelayer/generated/cosmosbridge"
	oracle "github.com/cosmos/peggy/cmd/ebrelayer/generated/oracle"
)

const (
	// GasLimit : the gas limit in Gwei used for transactions sent with TransactOpts
	GasLimit = uint64(300000)
)

// RelayProphecyClaimToEthereum :
func RelayProphecyClaimToEthereum(provider string, contractAddress common.Address, event events.Event, msgData events.CosmosMsg) error {
	client, auth, target := getRelayConfig(provider, contractAddress, event)

	fmt.Println("\nFetching CosmosBridge contract...")

	// Initialize CosmosBridge instance
	cosmosBridgeInstance, err := cosmosBridge.NewCosmosBridge(target, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ClaimType:", uint8(msgData.ClaimType))

	fmt.Println("\nmsgData:", msgData)

	fmt.Println("Sending new ProphecyClaim to CosmosBridge...")
	tx, err := cosmosBridgeInstance.NewProphecyClaim(auth, uint8(msgData.ClaimType), msgData.CosmosSender, msgData.EthereumReceiver, msgData.TokenContractAddress, msgData.Symbol, msgData.Amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nNewProphecyClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// Report tx status
	switch receipt.Status {
	case 0:
		fmt.Println("Status: 0 - Failed")
	case 1:
		fmt.Println("Status: 1 - Successful")
	}

	return nil
}

// RelayOracleClaimToEthereum :
func RelayOracleClaimToEthereum(provider string, contractAddress common.Address, event events.Event, prophecyClaim events.NewProphecyClaimEvent) error {
	oracleClaim := ProphecyClaimToOracleClaim(prophecyClaim)

	client, auth, target := getRelayConfig(provider, contractAddress, event)

	fmt.Println("\nFetching Oracle contract...")

	// Initialize Oracle instance
	oracleInstance, err := oracle.NewOracle(target, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sending new OracleClaim to Oracle...")

	tx, err := oracleInstance.NewOracleClaim(auth, oracleClaim.ProphecyID, oracleClaim.Message, oracleClaim.Signature)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nNewOracleClaim tx hash:", tx.Hash().Hex())
	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// Report tx status
	switch receipt.Status {
	case 0:
		fmt.Println("Status: 0 - Failed")
	case 1:
		fmt.Println("Status: 1 - Successful")
	}

	return nil
}

func getRelayConfig(provider string, registry common.Address, event events.Event) (client *ethclient.Client, auth *bind.TransactOpts, target common.Address) {
	// Start Ethereum client
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	// Load the validator's private key
	key, err := LoadPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Load the validator's address
	sender, err := LoadSender()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Set up TransactOpts auth's tx signature authorization
	transactOptsAuth := bind.NewKeyedTransactor(key)
	transactOptsAuth.Nonce = big.NewInt(int64(nonce))
	transactOptsAuth.Value = big.NewInt(0) // in wei
	transactOptsAuth.GasLimit = GasLimit
	transactOptsAuth.GasPrice = gasPrice

	// All ProphecyClaims are made to the CosmosBridge contract
	var targetContract ContractRegistry

	switch event {
	case events.MsgBurn, events.MsgLock:
		targetContract = CosmosBridge
	case events.LogNewProphecyClaim:
		targetContract = Oracle
	}

	// Get the specific contract's address
	target, err = GetAddressFromBridgeRegistry(client, registry, targetContract)
	if err != nil {
		log.Fatal(err)
	}

	return client, auth, target
}
