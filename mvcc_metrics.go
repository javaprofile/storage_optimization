package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type Version struct {
	Timestamp int64
	Value     string
}

type MVCCStore struct {
	data   map[string][]Version
	mutex  sync.RWMutex
}

func NewMVCCStore() *MVCCStore {
	return &MVCCStore{
		data: make(map[string][]Version),
	}
}

func (s *MVCCStore) Write(key string, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	timestamp := time.Now().UnixNano()
	s.data[key] = append(s.data[key], Version{Timestamp: timestamp, Value: value})
}

func (s *MVCCStore) Read(key string, timestamp int64) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	versions := s.data[key]
	for i := len(versions) - 1; i >= 0; i-- {
		if versions[i].Timestamp <= timestamp {
			return versions[i].Value, true
		}
	}
	return "", false
}

type StorageMetrics struct {
	TotalKeys     int
	TotalVersions int
	ApproxMemBytes int
}

func (s *MVCCStore) CollectMetrics() StorageMetrics {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var keys, versions, mem int
	for k, vList := range s.data {
		keys++
		mem += len(k)
		for _, v := range vList {
			versions++
			mem += len(v.Value) + int(unsafe.Sizeof(v.Timestamp))
		}
	}

	return StorageMetrics{
		TotalKeys:     keys,
		TotalVersions: versions,
		ApproxMemBytes: mem,
	}
}

func main() {
	store := NewMVCCStore()
	store.Write("x", "value1")
	time.Sleep(time.Millisecond)
	store.Write("x", "value2")
	store.Write("y", "valA")

	metrics := store.CollectMetrics()
	fmt.Printf("Keys: %d, Versions: %d, ApproxMemory(bytes): %d\n", metrics.TotalKeys, metrics.TotalVersions, metrics.ApproxMemBytes)
}
