package heap

type KthLargest struct {
	h *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	x := NewIntHeap(k)
	for _, n := range nums {
		x.Insert(n)
	}
	r := KthLargest{
		h: x,
	}
	return r
}

func (this *KthLargest) Add(val int) int {
	this.h.Insert(val)
	return this.h.Min()
}
