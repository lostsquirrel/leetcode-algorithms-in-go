package queue

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 || len(nums) == 0 {
		return nums
	}
	r := make([]int, len(nums)-k+1)
	w := &IntDeque{}
	rindex := 0

	for i, v := range nums {
		if i >= k && w.PeekFront() <= i-k {
			w.PopFront()
		}
		for !w.Empty() && v >= nums[w.PeekBack()] {
			w.PopBack()
		}
		w.PushBack(i)

		if i >= k-1 {
			r[rindex] = nums[w.PeekFront()]
			rindex += 1
		}
	}
	return r
}
