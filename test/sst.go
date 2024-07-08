package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/node"
)

func main() {
    // 노드 구성 설정
    nodeConfig := node.DefaultConfig
    nodeConfig.DataDir = "/mnt/nvme0n1/ehtereum/execution/Fullnode/geth/"
    nodeConfig.Name = "geth"

    // 노드 인스턴스 생성
    n, err := node.New(&nodeConfig)
    if err != nil {
        log.Fatalf("Failed to create the protocol stack: %v", err)
    }

    // Geth 데이터베이스 열기
    db, err := n.OpenDatabase("chaindata", 0, 0, "",false)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    // 키-값 데이터 순회 예제
    iter := db.NewIterator(nil,nil)
    for iter.Next() {
        key := iter.Key()
        value := iter.Value()
        fmt.Printf("Key: %s, Value: %x\n", key, value)
    }
    iter.Release()
    if err := iter.Error(); err != nil {
        log.Fatalf("Iterator error: %v", err)
    }
}

func formatHex(data []byte) string {
    return hex.EncodeToString(data)
}