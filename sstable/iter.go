package main

import (
    "encoding/hex"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/node"
)

func main() {
    // 노드 구성 설정
    nodeConfig := node.DefaultConfig
    nodeConfig.DataDir = "/mnt/nvme0n1/ethereum/execution/Fullnode/geth/"
    nodeConfig.Name = "geth"

    // 노드 인스턴스 생성
    n, err := node.New(&nodeConfig)
    if err != nil {
        log.Fatalf("Failed to create the protocol stack: %v", err)
    }

    // Geth 데이터베이스 열기
    db, err := n.OpenDatabase("chaindata", 0, 0, "", false)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()


    // 특정 키 검색 (예시)
    searchKey, _ := hex.DecodeString("b28471d6118a65721c8c1a01eb190aa32eda805e5daf8ace7499dc03cf4b80bc")
    value, err := db.Get(searchKey)
    if err != nil {
        fmt.Printf("Key not found: %s\n", formatHex(searchKey))
    } else {
        fmt.Printf("Found value for key %s: %s\n", formatHex(searchKey), formatHex(value))
    }
}

func formatHex(data []byte) string {
    return hex.EncodeToString(data)
}