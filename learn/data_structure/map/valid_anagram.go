package map_demo

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		_, ok := m[r]
		if !ok {
			m[r] = 1
		} else {
			m[r] += 1
		}
	}
	for _, r := range t {
		n, ok := m[r]
		if !ok {
			return false
		}
		if n < 1 {
			return false
		}
		m[r] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
