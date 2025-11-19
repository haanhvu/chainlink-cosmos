package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	// import your generated bindings package
	"github.com/haanhvu/chainlink-polkadot/relayer/build/bindings"
	//bolt "go.etcd.io/bbolt"
)

var (
	ethWSUrl        = os.Getenv("ETH_WS")  // e.g. wss://sepolia.infura.io/ws/v3/KEY
	ethRPC          = os.Getenv("ETH_RPC") // fallback http RPC
	consumerAddress = common.HexToAddress(os.Getenv("CONSUMER_ADDR"))
	//substrateUrl    = os.Getenv("SUBSTRATE_WS")   // wss://
	//dbPath = os.Getenv("DB_PATH") // ./relayer.db
	//substrateSeed   = os.Getenv("SUBSTRATE_SEED") // mnemonic or raw secret
)

//const processedBucket = "processed"

func main() {
	// open local DB for processed IDs
	/*db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(processedBucket))
		return err
	})*/

	// connect to eth (prefer ws)
	ethClient, err := ethclient.Dial(ethWSUrl)
	if err != nil {
		log.Printf("ws dial failed, fallback to http: %v", err)
		ethClient, err = ethclient.Dial(ethRPC)
		if err != nil {
			log.Fatalf("can't dial eth: %v", err)
		}
	}
	defer ethClient.Close()

	/*// connect substrate
	  subAPI, err := gsrpc.NewSubstrateAPI(substrateUrl)
	  if err != nil {
	      log.Fatalf("sub api: %v", err)
	  }
	  meta, err := subAPI.RPC.State.GetMetadataLatest()
	  if err != nil {
	      log.Fatalf("get metadata: %v", err)
	  }
	  _ = meta // may be used for call construction*/

	// create contract instance
	contract, err := bindings.NewFunctionsConsumer(consumerAddress, ethClient)
	if err != nil {
		log.Fatalf("new contract: %v", err)
	}

	// context + graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	// subscribe to logs using binding's Watch... or Filter...; here we use a Filter and parse logs manually
	query := ethereum.FilterQuery{
		Addresses: []common.Address{consumerAddress},
	}
	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("subscribe logs: %v", err)
	}
	log.Println("listening for events...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("eth subscription error: %v", err)
			// try to reconnect / exit
			time.Sleep(2 * time.Second)
			// TODO: robust reconnect logic
		case vLog := <-logs:
			// Try to parse using generated binding event parser
			ev, parseErr := contract.ParseFunctionsDataUpdated(vLog)
			if parseErr != nil {
				log.Printf("parse event err: %v", parseErr)
				continue
			}
			// requestId is [32]byte
			requestId := ev.RequestId
			// response bytes
			resp := ev.Response
			// err bytes
			errBytes := ev.Error

			// idempotency: check db
			/*already := false
			db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(processedBucket))
				if b.Get(requestId[:]) != nil {
					already = true
				}
				return nil
			})
			if already {
				log.Printf("request %x already processed, skipping", requestId)
				continue
			}*/

			// process: decode transform as needed
			log.Printf("got event requestId=%x response=%s err=%x", requestId, string(resp), errBytes)

			// submit to substrate (with retry/backoff)
			op := func() error {
				return nil
				//return submitToSubstrate(subAPI, requestId, resp)
			}
			bo := backoff.NewExponentialBackOff()
			bo.MaxElapsedTime = 2 * time.Minute
			if err := backoff.Retry(op, bo); err != nil {
				log.Printf("failed submit after retries: %v", err)
				continue
			}

			// mark processed
			/*db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(processedBucket))
				return b.Put(requestId[:], []byte(time.Now().Format(time.RFC3339)))
			})*/

		case <-ctx.Done():
			log.Println("shutting down")
			return
		}
	}
}

/*// submitToSubstrate constructs and sends extrinsic; adapt pallet/call names & types
func submitToSubstrate(api *gsrpc.SubstrateAPI, requestId [32]byte, data []byte) error {
    // example: call pallet 'oracle' with function 'set_sentiment'
    meta, err := api.RPC.State.GetMetadataLatest()
    if err != nil {
        return fmt.Errorf("metadata: %v", err)
    }

    // create call: module 'OraclePallet' and call 'set_sentiment'; change to your pallet/module names
    // Assume function signature: fn set_sentiment(origin, request_id: [u8;32], data: Vec<u8>)
    c, err := types.NewCall(meta, "OracleModule.set_sentiment", requestId, data)
    if err != nil {
        return fmt.Errorf("new call: %v", err)
    }

    // load signer keypair from seed
    kp, err := signature.KeyringPairFromSecret(substrateSeed, 42) // 42 = SS58 prefix; change to your chain
    if err != nil {
        return fmt.Errorf("keypair: %v", err)
    }

    // build extrinsic
    ext := types.NewExtrinsic(c)
    // get nonce
    genesisHash, err := api.RPC.Chain.GetBlockHash(0)
    if err != nil {
        return fmt.Errorf("get genesis: %v", err)
    }
    runtimeVersion, err := api.RPC.State.GetRuntimeVersionLatest()
    if err != nil {
        return fmt.Errorf("runtime ver: %v", err)
    }
    key, err := types.CreateStorageKey(meta, "System", "Account", kp.PublicKey, nil)
    if err != nil {
        return fmt.Errorf("create storage key: %v", err)
    }
    var accountInfo types.AccountInfo
    ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
    if err != nil {
        return fmt.Errorf("get account: %v", err)
    }
    var nonce uint32
    if ok {
        nonce = uint32(accountInfo.Nonce)
    } else {
        nonce = 0
    }

    // construct signature payload
    o := types.SignatureOptions{
        BlockHash:          genesisHash,
        Era:                types.ExtrinsicEra{IsMortalEra: false},
        GenesisHash:        genesisHash,
        Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
        SpecVersion:        runtimeVersion.SpecVersion,
        TransactionVersion: runtimeVersion.TransactionVersion,
        Tip:                types.NewUCompactFromUInt(0),
    }

    // sign the extrinsic
    if err := ext.Sign(kp, o); err != nil {
        return fmt.Errorf("sign ext: %v", err)
    }

    // submit
    sub, err := api.RPC.Author.SubmitExtrinsic(ext)
    if err != nil {
        return fmt.Errorf("submit: %v", err)
    }
    log.Printf("submitted extrinsic: %v", sub)
    return nil
}*/
