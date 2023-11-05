package hashring

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var nodes = []string{
	"app-d6955dfbc-6gzsd",
	"app-d6955dfbc-sh6tl",
	"app-d6955dfbc-5dfbc",
}

func TestRingHashing(t *testing.T) {
	ring := NewRing(1, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	ring.Add("2", "4")
	assert.Equal(t, []uint32{2, 4}, ring.keys)

	ring.Add("6", "8", "10")
	assert.Equal(t, []uint32{2, 4, 6, 8, 10}, ring.keys)

	ring.Remove("6", "8", "10")
	assert.Equal(t, []uint32{2, 4}, ring.keys)
}

func TestRingConsistency(t *testing.T) {
	ring1 := NewRing(3, nil)
	ring2 := NewRing(3, nil)

	ring1.Add(nodes...)
	ring2.Add(nodes...)

	for tt := range nodes {
		assert.Equal(t, ring1.Get(nodes[tt]), ring2.Get(nodes[tt]))
	}
}

func TestRingMine(t *testing.T) {
	ring := NewRing(3, nil)
	ring.Add(nodes...)

	tests := []struct {
		key      string
		name     string
		expected bool
	}{
		{"56e7af44-c781-4d39-9244-05ae7601b5a4", "app-d6955dfbc-sh6tl", true},
		{"16486cc5-6212-4f08-ab5b-b175d162fb90", "app-d6955dfbc-5dfbc", true},
		{"64a9c5cb-c103-4e6e-8edd-4668d5c45e53", "app-d6955dfbc-6gzsd", true},
		{"64a9c5cb-c103-4e6e-8edd-4668d5c45e53", "app-d6955dfbc-5dfbc", false},
	}

	for tt := range tests {
		assert.Equal(t, tests[tt].expected, ring.Mine(tests[tt].name, tests[tt].key))
	}
}
