package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/node"
    "github.com/ethereum/go-ethereum/core/rawdb"
)

const (
    // 노드 구성 설정
    cacheSize = 1024
	handles   = 100
)
func main(){

    //노드 세부 설정
    nodeConfig := node.DefaultConfig
    nodeConfig.DataDir = "/mnt/nvme0n1/ethereum/execution/Fullnode/geth/"
    nodeConfig.Name = "geth"

    // 노드 인스턴스 생성
    n, err := node.New(&nodeConfig)
    if err != nil {
        log.Fatalf("Failed to create the protocol stack: %v", err)
    }

    // Geth 데이터베이스 열기
    db, err := n.OpenDatabase("chaindata", cacheSize, handles, "", false)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()
    log.Println("Database opened successfully")
    
    // Ancient 데이터베이스 오픈
    ancient,err := rawdb.NewDatabaseWithFreezer(db, nodeConfig.DataDir+"/chaindata/ancient", "",false)
    if err != nil{
        log.Fatalf("failed to open ancient")
    }

    // 최신 블록 번호 가져오기
    latestBlock := rawdb.ReadHeaderNumber(ancient, rawdb.ReadHeadHeaderHash(ancient))
    if latestBlock == nil {
        log.Fatalf("failed to read block number")
    }

    fmt.Printf("Latest block number: %d", *latestBlock)
}