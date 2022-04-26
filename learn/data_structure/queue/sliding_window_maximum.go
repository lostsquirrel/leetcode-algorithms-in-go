package queue

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	r := make([]int, len(nums)-k+1)
	rindex := 0
	var m1 *MaxNode
	var m2 *MaxNode
	head := 0
	// tail := k

	for i := 0; i < k; i++ {
		e := nums[i]
		if m1 == nil {
			m1 = &MaxNode{i, e}
		} else {
			if e > m1.value {
				m1.index = i
				m1.value = e
			} else {
				if m2 == nil {
					m2 = &MaxNode{i, e}
				} else {
					if e > m2.value {
						m2.index = i
						m2.value = e
					}
				}
			}
		}

	}
	r[rindex] = m1.value
	for i := k; i < len(nums); i++ {
		e := nums[i]
		head += 1
		if m1.index < head {
			// m1 should gone
			if e > m2.value {
				m1.index = i
				m1.value = e
				m2 = nil
			} else {
				m1.index = i
				m1.value = e
				m1, m2 = m2, m1
			}
		} else {
			if e > m1.value {
				m1.index = i
				m1.value = e
				// m1, m2 = m2, m1
			} else {
				if m2 == nil {
					m2 = &MaxNode{i, e}
				} else {
					if e > m2.value {
						m2.index = i
						m2.value = e
					}
				}
			}
		}
		rindex += 1
		r[rindex] = m1.value

	}
	// r[rindex] = m1.value
	return r
}

type MaxNode struct {
	index int
	value int
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
