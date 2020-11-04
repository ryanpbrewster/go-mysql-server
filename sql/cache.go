package sql

import (
	"fmt"
	"github.com/cespare/xxhash"
	"runtime"

	lru "github.com/hashicorp/golang-lru"
	errors "gopkg.in/src-d/go-errors.v1"
)

// CacheKey returns a hash of the given value to be used as key in
// a cache.
func CacheKey(v interface{}) (uint64, error) {
	hash := xxhash.New()
	if _, err := hash.Write([]byte(fmt.Sprintf("%#v", v))); err != nil {
		return 0, err
	}

	return hash.Sum64(), nil
}

// ErrKeyNotFound is returned when the key could not be found in the cache.
var ErrKeyNotFound = errors.NewKind("memory: key %d not found in cache")

type lruCache struct {
	memory   Freeable
	reporter Reporter
	size     int
	cache    *lru.Cache
}

func (l *lruCache) Size() int {
	return l.size
}

func newLRUCache(memory Freeable, r Reporter, size uint) *lruCache {
	lru, _ := lru.New(int(size))
	return &lruCache{memory, r, int(size), lru}
}

func (l *lruCache) Put(k uint64, v interface{}) error {
	if releaseMemoryIfNeeded(l.reporter, l.Free, l.memory.Free) {
		l.cache.Add(k, v)
	}
	return nil
}

func (l *lruCache) Get(k uint64) (interface{}, error) {
	v, ok := l.cache.Get(k)
	if !ok {
		return nil, ErrKeyNotFound.New(k)
	}

	return v, nil
}

func (l *lruCache) Free() {
	l.cache, _ = lru.New(l.size)
}

func (l *lruCache) Dispose() {
	l.memory = nil
	l.cache = nil
}

type rowsCache struct {
	memory   Freeable
	reporter Reporter
	rows     []Row
}

func newRowsCache(memory Freeable, r Reporter) *rowsCache {
	return &rowsCache{memory, r, nil}
}

func (c *rowsCache) Add(row Row) error {
	if !releaseMemoryIfNeeded(c.reporter, c.memory.Free) {
		return ErrNoMemoryAvailable.New()
	}

	c.rows = append(c.rows, row)
	return nil
}

func (c *rowsCache) Get() []Row { return c.rows }

func (c *rowsCache) Dispose() {
	c.memory = nil
	c.rows = nil
}

// mapCache is a simple in-memory implementation of a cache
type mapCache struct {
	cache map[uint64]interface{}
}

func (m mapCache) Put(u uint64, i interface{}) error {
	m.cache[u] = i
	return nil
}

func (m mapCache) Get(u uint64) (interface{}, error) {
	return m.cache[u], nil
}

func (m mapCache) Size() int {
	return len(m.cache)
}

func NewMapCache() mapCache {
	return mapCache{
		cache: make(map[uint64]interface{}),
	}
}

type historyCache struct {
	memory   Freeable
	reporter Reporter
	cache    map[uint64]interface{}
}

func (h *historyCache) Size() int {
	return len(h.cache)
}

func newHistoryCache(memory Freeable, r Reporter) *historyCache {
	return &historyCache{memory, r, make(map[uint64]interface{})}
}

func (h *historyCache) Put(k uint64, v interface{}) error {
	if !releaseMemoryIfNeeded(h.reporter, h.memory.Free) {
		return ErrNoMemoryAvailable.New()
	}
	h.cache[k] = v
	return nil
}

func (h *historyCache) Get(k uint64) (interface{}, error) {
	v, ok := h.cache[k]
	if !ok {
		return nil, ErrKeyNotFound.New(k)
	}
	return v, nil
}

func (h *historyCache) Dispose() {
	h.memory = nil
	h.cache = nil
}

// releasesMemoryIfNeeded releases memory if needed using the following steps
// until there is available memory. It returns whether or not there was
// available memory after all the steps.
func releaseMemoryIfNeeded(r Reporter, steps ...func()) bool {
	for _, s := range steps {
		if HasAvailableMemory(r) {
			return true
		}

		s()
		runtime.GC()
	}

	return HasAvailableMemory(r)
}
