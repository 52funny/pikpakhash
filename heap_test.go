package pikpakhash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	hash := make([]segmentHash, 0)
	hash1 := segmentHash{id: 1, hash: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	hash2 := segmentHash{id: 2, hash: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	hash3 := segmentHash{id: 3, hash: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	hash4 := segmentHash{id: 4, hash: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	hash = append(hash, hash1)
	hash = append(hash, hash4)
	hash = append(hash, hash3)
	hash = append(hash, hash2)
	h := new(heap)
	for _, v := range hash {
		h.Push(v)
	}
	assert.Equal(t, h.Pop(), hash1)
	assert.Equal(t, h.Pop(), hash2)
	assert.Equal(t, h.Pop(), hash3)
	assert.Equal(t, h.Peek(), &hash4)
	assert.Equal(t, h.Pop(), hash4)
}
