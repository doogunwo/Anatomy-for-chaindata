package main
import (
	
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"

)

func getLDB(ldbPath string) ethdb.Database {
	cache := 256
	handles := 256
	ldb, err := rawdb.NewLevelDBDatabase(ldbPath, cache, handles, "", true)
	if err != nil {
		fmt.Println("Did not find leveldb at path:", ldbPath)
		fmt.Println("Are you sure you are pointing to the 'chaindata' folder?")
		panic(err)
	}
	fmt.Print("LevelDB ok\n")
	return ldb
}

func main(){
	chaindataPath := "/mnt/nvme0n1/ehtereum/execution/Fullnode/geth/chaindata"
	db, err := leveldb.OpenFile(path, &opt.Options{ReadOnly: true})
    if err != nil {
        panic(err)
    }
    defer db.Close()
}