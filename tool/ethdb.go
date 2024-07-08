package main

import (
    "log"

    "github.com/ethereum/go-ethereum/ethdb"
)


func main() {
    db, err := ethdb.NewLDBDatabase("/path/to/your/.ethereum/geth/chaindata", 0, 0)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    // Your code to interact with the database goes here
}

