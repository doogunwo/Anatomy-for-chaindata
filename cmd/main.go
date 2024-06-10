package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rcrowley/go-metrics"
	"github.com/ethereum/go-ethereum/ethdb"
    
	

	"fmt"
	"log"

)

const (
	ipcPath = "/mnt/nvme0n1/ehtereum/execution/node160/geth.ipc"
	chaindataPath ="/mnt/nvme0n1/ehtereum/execution/node160/geth/chaindata"
)

func main(){

	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatalf("failed to connect th the Ethereum client : %v", err)
	}
	defer client.Close()
	

	registry := metrics.NewRegistry()

	// leveldb Get 함수 호출 횟수 메트릭 등록
	getCounter := metrics.NewCounter()
	registry.Register("leveldb.gets", getCounter)

	
}

