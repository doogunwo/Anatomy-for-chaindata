package main

import (
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/ethdb/leveldb"
)

func main() {
    // 데이터베이스 경로 지정
    dbPath := "/mnt/nvme0n1/ehtereum/execution/Fullnode/geth/chaindata"

    // ethdb를 사용하여 LevelDB 열기
    db, err := leveldb.New(dbPath, 1024, 100, "", false)
    if err != nil {
        log.Fatalf("Failed to open LevelDB with ethdb: %v", err)
    } else {
        log.Println("Database opened successfully")
    }
    defer db.Close()

    // ethdb를 사용하여 데이터베이스 순회
    it := db.NewIterator(nil, nil)
    if it == nil {
        log.Println("Iterator is nil")
    } else {
        log.Println("Iterator created successfully")
    }

    found := false
    for it.Next() {
        found = true
        key := it.Key()
        value := it.Value()
        fmt.Printf("Key: %x, Value: %x\n", key, value)
    }
    if !found {
        log.Println("No entries found in the database")
    }
    it.Release()
    if err := it.Error(); err != nil {
        log.Fatalf("Iterator error: %v", err)
    }
}

