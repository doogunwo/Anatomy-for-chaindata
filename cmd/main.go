package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/doogunwo/OCGethStorage/pkg/ethclient"
)

func main() {
    ipcPath := flag.String("ipcPath", "", "Path to geth.ipc")
    flag.Parse()

    if *ipcPath == "" {
        log.Fatal("ipcPath is required")
    }

    client, err := ethclient.Dial(*ipcPath)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Example usage
    blockNumber, err := client.BlockNumber()
    if err != nil {
        log.Fatalf("Failed to retrieve block number: %v", err)
    }

    fmt.Printf("Current block number: %d\n", blockNumber)
}
