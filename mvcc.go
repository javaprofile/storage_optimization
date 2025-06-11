package main
import (
	"fmt"
	"sort"
	"sync"
	"time"
)
type Version struct {
	Timestamp int64
	Value     string
}
type MVCCStore struct {
	mu    sync.RWMutex
	store map[string][]Version
}
func NewMVCCStore() *MVCCStore {
	return &MVCCStore{
		store: make(map[string][]Version),
	}
}
func (mvcc *MVCCStore) Write(key, value string, timestamp int64) {
	mvcc.mu.Lock()
	defer mvcc.mu.Unlock()
	mvcc.store[key] = append(mvcc.store[key], Version{
		Timestamp: timestamp,
		Value:     value,
	})
}
func (mvcc *MVCCStore) Read(key string, timestamp int64) (string, bool) {
	mvcc.mu.RLock()
	defer mvcc.mu.RUnlock()
	versions := mvcc.store[key]
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Timestamp > versions[j].Timestamp
	})
	for _, v := range versions {
		if v.Timestamp <= timestamp {
			return v.Value, true
		}
	}
	return "", false
}
func now() int64 {
	return time.Now().UnixNano()
}
func main() {
	mvcc := NewMVCCStore()
	t1 := now()
	mvcc.Write("user:1", "Alice", t1)
	time.Sleep(time.Millisecond * 10)
	t2 := now()
	mvcc.Write("user:1", "Bob", t2)
	time.Sleep(time.Millisecond * 10)
	t3 := t1 + 5
	value, ok := mvcc.Read("user:1", t3)
	if ok {
		fmt.Println("Read at t3:", value)
	}
	t4 := now()
	value2, ok2 := mvcc.Read("user:1", t4)
	if ok2 {
		fmt.Println("Read at t4:", value2)
	}
}
