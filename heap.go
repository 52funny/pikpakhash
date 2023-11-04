package pikpakhash

// binary heap definition
type heap struct {
	data []segmentHash
}

func newHeap() heap {
	return heap{
		data: make([]segmentHash, 0),
	}
}

// return the length of the inner slice
func (h *heap) Len() int {
	return len(h.data)
}

// up heapify
func (h *heap) up(pos int) {
	for pos > 0 && h.data[(pos-1)/2].id > h.data[pos].id {
		father := (pos - 1) / 2
		h.data[father], h.data[pos] = h.data[pos], h.data[father]
		pos = father
	}
}

// down heapify
func (h *heap) down(pos int) {
	n := h.Len()
	for pos*2+1 < n {
		min := pos*2 + 1
		right := pos*2 + 2
		if right < n && h.data[min].id > h.data[right].id {
			min = right
		}
		if h.data[pos].id > h.data[min].id {
			h.data[pos], h.data[min] = h.data[min], h.data[pos]
		}
		pos = min
	}
}

// push a new element into the heap
func (h *heap) Push(val segmentHash) {
	h.data = append(h.data, val)
	h.up(h.Len() - 1)
}

// pop the top element from the heap
// if heap inner data size is zero will be panic
func (h *heap) Pop() segmentHash {
	n := h.Len() - 1
	if n < 0 {
		panic("heap is empty")
	}
	h.data[0], h.data[n] = h.data[n], h.data[0]
	val := h.data[n]
	h.data = h.data[:n]
	h.down(0)
	return val
}

// get the top element pointer from the heap
// if heap inner data size is zero will return nil
func (h *heap) Peek() *segmentHash {
	n := h.Len() - 1
	if n < 0 {
		return nil
	}
	return &h.data[0]
}
