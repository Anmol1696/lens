package main

import (
	"os"

	"go.uber.org/zap"

	lens "github.com/strangelove-ventures/lens/client"
	"github.com/strangelove-ventures/lens/client/query"
)

func NewClient(logger *zap.Logger) *lens.ChainClient {
	ccc := &lens.ChainClientConfig{
		ChainID:        "cosmoshub-4",
		RPCAddr:        "https://cosmos-rpc.polkachu.com:443",
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "30s",
		SignModeStr:    "direct",
		Modules:        lens.ModuleBasics,
	}
	client, err := lens.NewChainClient(logger, ccc, os.Getenv("HOME"), os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	// Setup logger, client and querier
	logger := zap.NewExample()
	client := NewClient(logger)
	querier := query.Query{Client: client, Options: query.DefaultOptions()}

	// Works well
	logger.Info("fetch particular channel")
	channel, err := querier.Ibc_Channel("channel-0", "transfer")
	if err != nil {
		panic(err)
	}
	logger.Info("channels fetched", zap.Any("channel", channel))

	// Does not work
	clientId := "07-tendermint-1"
	logger.Info("going to make request for client state", zap.String("client-id", clientId))
	state, err := querier.Ibc_ClientState(clientId)
	if err != nil {
		panic(err)
	}
	logger.Info("client state retrived", zap.Any("state", state))
}
