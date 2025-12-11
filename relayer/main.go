package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/haanhvu/chainlink-cosmos/relayer/bindings"
)

var (
	ethWSUrl        = os.Getenv("ETH_WS")
	ethRPC          = os.Getenv("ETH_RPC")
	consumerAddress = common.HexToAddress(os.Getenv("CONSUMER_ADDR"))
)

func main() {
	ethClient, err := ethclient.Dial(ethWSUrl)
	if err != nil {
		log.Printf("ws dial failed, fallback to http: %v", err)
		ethClient, err = ethclient.Dial(ethRPC)
		if err != nil {
			log.Fatalf("can't dial eth: %v", err)
		}
	}
	defer ethClient.Close()

	filterer, err := bindings.NewFunctionsConsumerFilterer(consumerAddress, ethClient)
	if err != nil {
		log.Fatalf("error creating log filterer for FunctionsConsumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{consumerAddress},
	}
	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("error subscribing filter logs: %v", err)
	}
	log.Println("listening for filter log events...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("eth subscription error: %v", err)
			// TODO: handle error
		case vLog := <-logs:
			log.Printf("Raw log topics: %v\nData: %x\n", vLog.Topics, vLog.Data)

			ev, parseErr := filterer.ParseFunctionsDataUpdated(vLog)
			if parseErr != nil {
				log.Printf("parse event err: %v", parseErr)
				continue
			}

			requestId := ev.RequestId
			resp := ev.Response
			errBytes := ev.Error

			log.Printf("got event requestId=%x response=%s err=%x", requestId, string(resp), string(errBytes))

			submitToCosmos(resp)

		case <-ctx.Done():
			log.Println("shutting down")
			return
		}
	}
}

func submitToCosmos(data []byte) error {
	// TODO: Implement:
	//	- Wrap data inside an transaction
	//	- Sign the transaction
	//	- Broadcast the transaction
	return nil
}
