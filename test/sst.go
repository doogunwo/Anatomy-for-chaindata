package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	// LevelDB 데이터베이스 경로 지정
    options := &opt.Options{
		BlockCacheCapacity: 10 * opt.MiB,     // 10MB의 블록 캐시
		WriteBuffer:        10 * opt.MiB,     // 10MB의 쓰기 버퍼
		OpenFilesCacheCapacity: 500,          // 열 수 있는 파일 핸들 수
	}

	path := "/mnt/nvme0n1/ehtereum/execution/archive40/geth/chaindata/"
	
	// 데이터베이스 열기
	db, err := leveldb.New(path, options)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}
	defer db.Close()

	// 데이터베이스에서 키-값 쌍 순회
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		log.Fatalf("Iterator error: %v", err)
	}
}
