package other

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		s := target - v
		for j := i + 1; j < len(nums); j++ {
			if s == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
