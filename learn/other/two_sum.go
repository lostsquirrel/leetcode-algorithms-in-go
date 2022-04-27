package other

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		s := target - v
		j, ok := m[s]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
