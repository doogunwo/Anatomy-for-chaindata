package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethdb/leveldb"
    "github.com/ethereum/go-ethereum/trie"
)

func main() {
    // LevelDB 열기
    db, err := leveldb.New("./testdb", 0, 0, "", false)
    if err != nil {
        log.Fatalf("Failed to open LevelDB: %v", err)
    }
    defer db.Close()

    // 트리 데이터베이스 생성
    trieDB := trie.NewDatabase(db)

    // 새로운 머클 패트리샤 트리 생성
    tr, err := trie.New(common.Hash{}, trieDB)
    if err != nil {
        log.Fatalf("Failed to create trie: %v", err)
    }

    // 키-값 쌍 삽입
    testData := map[string]string{
        "key1": "value1",
        "key2": "value2",
        "key3": "value3",
    }

    for k, v := range testData {
        err = tr.Update([]byte(k), []byte(v))
        if err != nil {
            log.Fatalf("Failed to update trie: %v", err)
        }
    }

    // 트리 커밋
    root, err := tr.Commit(nil)
    if err != nil {
        log.Fatalf("Failed to commit trie: %v", err)
    }

    // 변경된 노드들을 데이터베이스에 저장
    if err := trieDB.Commit(root, false); err != nil {
        log.Fatalf("Failed to commit nodes to database: %v", err)
    }

    fmt.Printf("Trie Root: %x\n", root)

    // 새로운 트리 인스턴스 생성 (커밋된 상태로부터)
    tr, err = trie.New(root, trieDB)
    if err != nil {
        log.Fatalf("Failed to create new trie from committed state: %v", err)
    }

    // 값 읽기
    fmt.Println("\nRetrieving values from trie:")
    for k := range testData {
        v, err := tr.Get([]byte(k))
        if err != nil {
            log.Fatalf("Failed to get value for key %s: %v", k, err)
        }
        fmt.Printf("Key: %s, Value: %s\n", k, string(v))
    }

    // LevelDB의 내용 직접 조회
    fmt.Println("\nDirect LevelDB content:")
    iter := db.NewIterator(nil, nil)
    for iter.Next() {
        fmt.Printf("Key: %x, Value: %x\n", iter.Key(), iter.Value())
    }
    iter.Release()
    if err := iter.Error(); err != nil {
        log.Fatalf("Iterator error: %v", err)
    }
}
