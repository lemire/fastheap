package fastheap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	h := newHeapuint32()
	h.add(1)
	h.poll()
}
func TestNewHeap(t *testing.T) {
	h := newHeapuint32()
	for j := 0; j < 8; j++ {
		h.add(rand(j))
	}
	val := h.poll()
	for j := 1; j < 128; j++ {
		newval := h.poll()
		if val < newval {
			t.Errorf("Failed")
		}
		val = newval
	}
}

func testGoHeap(t *testing.T) {
	pq := make(PriorityQueue, 0)
	for j := 0; j < 128; j++ {
		heap.Push(&pq, rand(j))
	}
	val := heap.Pop(&pq).(uint32)
	for j := 1; j < 128; j++ {
		newval := heap.Pop(&pq).(uint32)
		if val < newval {
			t.Errorf("Failed")
		}
		val = newval
	}
}

// very fast semi-random function
func rand(i int) uint32 {
	i = i + 10000
	i = i ^ (i << 16)
	i = (i >> 5) ^ i
	return uint32(i & 0xFF)
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := newHeapuint32()
		for j := 0; j < 128; j++ {
			h.add(rand(j))
		}
		for j := 0; j < 128*10; j++ {
			h.add(rand(j))
			h.poll()
		}
	}
}

func BenchmarkGoHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pq := make(PriorityQueue, 0)
		for j := 0; j < 128; j++ {
			heap.Push(&pq, rand(j))
		}
		for j := 0; j < 128*10; j++ {
			heap.Push(&pq, rand(j))
			heap.Pop(&pq)
		}
	}
}
