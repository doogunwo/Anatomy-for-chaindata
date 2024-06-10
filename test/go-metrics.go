package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rcrowley/go-metrics"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/metrics"
	

	"context"
	"fmt"
	"log"
	"time"

)

const (
	ipcPath := "/mnt/nvme0n1/ehtereum/execution/node160/geth.ipc"
	chaindataPath :="/mnt/nvme0n1/ehtereum/execution/node160/geth/chaindata"
)

func func3(){


	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatalf("failed to connect th the Ethereum client : %v", err)
	}
	defer client.Close()
	

	registry := metrics.NewRegistry()

	gets_counter := metrics.NewCounter()
	registry.Register("leveldb.gets".gets_counter)

	Gets := ethdb.NewLDBDatabase(chaindataPath,128,1024).Get

	ethdb.NewLDBDatabase(chaindataPath, 128, 1024).Get = func(key []byte) ([]byte, error) {
		getCounter.Inc(1)
		return originalGet(key)
	}

	for {
		fmt.Println(gets_counter.Count())
		time.sleep(1 * time.Second)
	}

}

