package main

import (
	"fmt"
	"sync"
	"time"
)

type StorageMetrics struct {
	TotalReads   int
	TotalWrites  int
	TotalStorage int64
	mu           sync.Mutex
}

func (sm *StorageMetrics) IncrementReads() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.TotalReads++
}

func (sm *StorageMetrics) IncrementWrites() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.TotalWrites++
}

func (sm *StorageMetrics) AddStorage(size int64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.TotalStorage += size
}

func (sm *StorageMetrics) PrintMetrics() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	fmt.Printf("Total Reads: %d\n", sm.TotalReads)
	fmt.Printf("Total Writes: %d\n", sm.TotalWrites)
	fmt.Printf("Total Storage: %d bytes\n", sm.TotalStorage)
}

func simulateStorageOperations(sm *StorageMetrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		sm.IncrementReads()
		time.Sleep(10 * time.Millisecond)
	}
}

func simulateWriteOperations(sm *StorageMetrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		sm.IncrementWrites()
		sm.AddStorage(100)
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	sm := &StorageMetrics{}

	wg.Add(2)
	go simulateStorageOperations(sm, &wg)
	go simulateWriteOperations(sm, &wg)

	wg.Wait()
	sm.PrintMetrics()
}
