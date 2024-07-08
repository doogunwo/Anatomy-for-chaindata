package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/ethdb/leveldb"
)

const (
	dbPath    = "/mnt/nvme0n1/ethereum/execution/Fullnode/geth/chaindata"
	cacheSize = 1024
	handles   = 100
)

func main() {
	// 탐색하고 싶은 key
	key := "0xb28471d6118a65721c8c1a01eb190aa32eda805e5daf8ace7499dc03cf4b80bc"

	// Open LevelDB using ethdb
	db, err := leveldb.New(dbPath, cacheSize, handles, "", false)
	if err != nil {
		log.Fatalf("Failed to open LevelDB with ethdb: %v", err)
	}
	log.Println("Database opened successfully")
	defer db.Close()

	// Analyze LevelDB
	levelTableCounts, hitProbabilities, err := analyzeLevelDB(db)
	if err != nil {
		log.Fatalf("Failed to analyze LevelDB: %v", err)
	}

	// Print results
	fmt.Println("Level Table Counts:", levelTableCounts)
	fmt.Println("Hit Probabilities:", hitProbabilities)

	
	// 특정 key 탐색
	searchKey(db, key)
}

func analyzeLevelDB(db *leveldb.Database) (map[int]int, map[int]float64, error) {
	levelTableCounts := make(map[int]int)
	hitProbabilities := make(map[int]float64)

	// Calculate table count for each level
	for level := 0; level <= 6; level++ {
		count, err := getTableCountForLevel(db, level)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting table count for level %d: %v", level, err)
		}
		levelTableCounts[level] = count
	}

	// Calculate hit probability for each level
	for level := 0; level <= 6; level++ {
		hitProbability, err := calculateHitProbability(db, level)
		if err != nil {
			return nil, nil, fmt.Errorf("error calculating hit probability for level %d: %v", level, err)
		}
		hitProbabilities[level] = hitProbability
	}

	return levelTableCounts, hitProbabilities, nil
}

func getTableCountForLevel(db *leveldb.Database, level int) (int, error) {
	property := fmt.Sprintf("leveldb.num-files-at-level%d", level)
	value, err := db.Stat(property)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

func calculateHitProbability(db *leveldb.Database, level int) (float64, error) {
	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	totalKeys := 0
	hitKeys := 0

	for iter.Next() {
		totalKeys++
		keyLevel, err := getKeyLevel(db, iter.Key())
		if err != nil {
			return 0, err
		}
		if keyLevel == level {
			hitKeys++
		}
	}

	if err := iter.Error(); err != nil {
		return 0, err
	}

	if totalKeys == 0 {
		return 0, nil
	}

	return float64(hitKeys) / float64(totalKeys), nil
}

func getKeyLevel(db *leveldb.Database, key []byte) (int, error) {
	property := "leveldb.sstables"
	value, err := db.Stat(property)
	if err != nil {
		return -1, err
	}

	return findLevelForKey(value, key)
}

func findLevelForKey(sstablesInfo string, key []byte) (int, error) {
	lines := strings.Split(sstablesInfo, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "L") {
			level, err := strconv.Atoi(line[1:2])
			if err != nil {
				return -1, fmt.Errorf("invalid level format: %s", line)
			}
			if containsSSTableForLevel(line, key) {
				return level, nil
			}
		}
	}
	return -1, fmt.Errorf("key not found in any level")
}

func containsSSTableForLevel(levelInfo string, key []byte) bool {
	parts := strings.Split(levelInfo, ":")
	if len(parts) < 2 {
		return false
	}
	tableInfo := parts[1]
	tables := strings.Split(tableInfo, ",")
	for _, table := range tables {
		if keyInRange(table, key) {
			return true
		}
	}
	return false
}

func keyInRange(tableRange string, key []byte) bool {
	parts := strings.Split(tableRange, "-")
	if len(parts) != 2 {
		return false
	}
	start, err := hex.DecodeString(strings.TrimSpace(parts[0]))
	if err != nil {
		return false
	}
	end, err := hex.DecodeString(strings.TrimSpace(parts[1]))
	if err != nil {
		return false
	}
	return bytes.Compare(key, start) >= 0 && bytes.Compare(key, end) <= 0
}

func saveResultsToCSV(filename string, data map[int]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for level, value := range data {
		if err := writer.Write([]string{strconv.Itoa(level), fmt.Sprintf("%v", value)}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}

	return nil
}

func searchKey(db *leveldb.Database, key string) {
	// Remove '0x' prefix if present
	if strings.HasPrefix(key, "0x") {
		key = key[2:]
	}

	// Convert hex string to byte slice
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		log.Fatalf("Failed to decode key: %v", err)
	}

	// Try to get the value for the key
	value, err := db.Get(keyBytes)
	if err != nil {
		log.Printf("Key not found or error occurred: %v", err)
	} else {
		fmt.Printf("Value for key %s: %x\n", key, value)
	}

	// Find the level of the key
	level, err := getKeyLevel(db, keyBytes)
	if err != nil {
		log.Printf("Failed to get key level: %v", err)
	} else {
		fmt.Printf("Key %s is at level: %d\n", key, level)
	}
}