package main

import (
	"fmt"
	"sync"
	"time"
)

type Record struct {
	ID      int
	Value   string
	Version int
}

type Database struct {
	records map[int]*Record
	mu      sync.Mutex
}

func NewDatabase() *Database {
	return &Database{
		records: make(map[int]*Record),
	}
}

func (db *Database) AddRecord(id int, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.records[id] = &Record{ID: id, Value: value, Version: 1}
}

func (db *Database) ReadRecord(id int) (*Record, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	record, exists := db.records[id]
	return record, exists
}

func (db *Database) UpdateRecord(id int, value string, version int) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	record, exists := db.records[id]
	if !exists || record.Version != version {
		return false
	}
	record.Value = value
	record.Version++
	return true
}

func main() {
	db := NewDatabase()
	db.AddRecord(1, "Initial Value")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		record, exists := db.ReadRecord(1)
		if exists {
			if db.UpdateRecord(1, "Updated by Client 1", record.Version) {
				fmt.Println("Client 1 updated the record")
			} else {
				fmt.Println("Client 1 failed due to version conflict")
			}
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		record, exists := db.ReadRecord(1)
		if exists {
			if db.UpdateRecord(1, "Updated by Client 2", record.Version) {
				fmt.Println("Client 2 updated the record")
			} else {
				fmt.Println("Client 2 failed due to version conflict")
			}
		}
	}()

	wg.Wait()
}
