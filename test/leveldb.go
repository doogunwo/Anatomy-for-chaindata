package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    chaindataPath := "/mnt/nvme0n1/ehtereum/execution/Fullnode/geth/chaindata"
    pattern := filepath.Join(chaindataPath, "*.sst")
    
    files, err := filepath.Glob(pattern)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        info, err := os.Stat(file)
        if err != nil {
            fmt.Printf("Error getting info for %s: %v\n", file, err)
            continue
        }
        fmt.Printf("File: %s, Size: %d bytes\n", file, info.Size())
    }
}
