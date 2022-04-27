package other

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		s := target - v
		j, ok := m[s]
		if ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}
