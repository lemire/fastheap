package fastheap

type heapuint32 struct {
	keys []uint32
}

func newHeapuint32() *heapuint32 {
	return &heapuint32{make([]uint32, 0, 16)}
}

func (h *heapuint32) add(myval uint32) {
	i := len(h.keys)
	h.keys = append(h.keys, myval)
	p := (i - 1) >> 1
	for (i > 0) && (myval < h.keys[p]) {
		h.keys[i] = h.keys[p]
		i = p
		p = (i - 1) >> 1
	}
	h.keys[i] = myval
}

//Look at the top of the queue (a smallest element)
func (h *heapuint32) peek() uint32 {
	return h.keys[0]
}

// remove the element on top of the heap (a smallest element)
func (h *heapuint32) poll() uint32 {
	ans := h.keys[0]
	if len(h.keys) > 1 {
		h.keys[0] = h.keys[len(h.keys)-1]
		h.keys = h.keys[:len(h.keys)-1]
		h._percolateDown(0)
	} else {
		h.keys = h.keys[:len(h.keys)-1]
	}
	return ans
}

func (h *heapuint32) _percolateDown(i int) {
	size := len(h.keys)
	ai := h.keys[i]
	l := (i << 1) + 1
	for l < size {
		i = l
		if l+1 < size {
			if h.keys[l+1] < h.keys[l] {
				i = l + 1
			}
		}
		h.keys[(i-1)>>1] = h.keys[i]
		l = (i << 1) + 1
	}
	p := (i - 1) >> 1
	for (i > 0) && (ai < h.keys[p]) {
		h.keys[i] = h.keys[p]
		i = p
		p = (i - 1) >> 1
	}
	h.keys[i] = ai
}

type PriorityQueue []uint32

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(uint32))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
