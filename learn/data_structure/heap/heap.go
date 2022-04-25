package heap

type IntHeap struct {
	data     []int
	capacity int
	length   int
}

func NewIntHeap(capacity int) *IntHeap {
	heap := &IntHeap{
		data:     make([]int, capacity+1),
		capacity: capacity,
		length:   0,
	}
	return heap
}

func (he *IntHeap) Insert(value int) {
	var i int
	if he.length < he.capacity {
		he.length += 1
		he.data[he.length] = value
		i = he.length

		for i/2 > 0 && he.data[i] < he.data[i/2] {
			he.swap(i, i/2)
			i = i / 2
		}
	} else {
		if value < he.data[1] {
			return
		}
		he.data[1] = value
		// he.swap(1, he.length)
		i = 1
		he.heapDown(1)
	}

}
func (he *IntHeap) heapUp(i int) {
	j := i / 2
	if j < 0 {
		return
	}
	if he.data[i] < he.data[i/2] {
		he.swap(i, j)
		he.heapUp(j)
	}
}

func (he *IntHeap) heapDown(i int) {
	j := i * 2
	if j > he.capacity {
		return
	}
	if he.data[i] > he.data[j] {
		he.swap(i, j)
		he.heapDown(j)
	}
	if j+1 > he.capacity {
		return
	}
	if he.data[i] > he.data[j+1] {
		he.swap(i, j+1)
		he.heapDown(j + 1)
	}

}

func (he *IntHeap) Min() int { return he.data[1] }

func (he *IntHeap) swap(i, j int) {
	he.data[i], he.data[j] = he.data[j], he.data[i]
}
