package queue

func maxSlidingWindow(nums []int, k int) []int {

	r := make([]int, len(nums)-k+1)
	rindex := 0

	for i := k; i <= len(nums); i++ {
		r[rindex] = findWindowMaxDouble(nums[rindex:i])
		rindex += 1
	}
	return r
}

type Window struct {
	max       int
	isHeadMax bool
}

func (w *Window) findWindowMax(nums []int) {

	for _, v := range nums {
		if v > w.max {
			w.max = v
			w.isHeadMax = false
		}
	}
}
func findWindowMaxDouble(nums []int) int {
	max := nums[0]
	i := 1
	j := len(nums) - 1
	for i <= j {
		if max < nums[i] {
			max = nums[i]
		}
		if max < nums[j] {
			max = nums[j]
		}
		i += 1
		j -= 1
	}
	return max
}
