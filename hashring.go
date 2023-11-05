package hashring

import (
	"cmp"
	"hash/crc32"
	"slices"
	"strconv"
	"sync"
)

type HashFn func([]byte) uint32

type Ring struct {
	fn       HashFn
	replicas int
	keys     []uint32
	nodes    map[uint32]string
	sync.RWMutex
}

func NewRing(replicas int, fn HashFn) *Ring {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}

	r := &Ring{
		nodes:    make(map[uint32]string),
		fn:       fn,
		replicas: replicas,
	}

	return r
}

func (r *Ring) IsEmpty() bool {
	r.RLock()
	defer r.RUnlock()

	return len(r.nodes) == 0
}

func (r *Ring) Add(keys ...string) {
	r.Lock()
	defer r.Unlock()

	for _, key := range keys {
		for i := 0; i < r.replicas; i++ {
			hash := uint32(r.fn([]byte(strconv.Itoa(i) + key)))
			r.keys = append(r.keys, hash)
			r.nodes[hash] = key
		}
	}

	r.sortKeys()
}

func (r *Ring) Remove(keys ...string) {
	r.Lock()
	defer r.Unlock()

	// TODO: Revisit this. It's not efficient.
	for _, key := range keys {
		for i := 0; i < r.replicas; i++ {
			hash := uint32(r.fn([]byte(strconv.Itoa(i) + key)))
			delete(r.nodes, hash)
			idx := slices.Index(r.keys, hash)
			if idx >= 0 {
				r.keys = slices.Delete(r.keys, idx, idx+1)
			}
		}
	}

	r.sortKeys()
}

func (r *Ring) Get(key string) string {
	if r.IsEmpty() {
		return ""
	}

	hash := r.fn([]byte(key))
	idx := slices.IndexFunc(r.keys, func(n uint32) bool {
		return n >= hash
	})

	if idx == -1 {
		idx = 0
	}

	return r.nodes[r.keys[idx]]
}

func (r *Ring) Mine(name, key string) bool {
	node := r.Get(key)
	return node == name
}

func (r *Ring) sortKeys() {
	slices.SortStableFunc(r.keys, func(a, b uint32) int {
		return cmp.Compare(a, b)
	})
}
