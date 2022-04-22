package stack

func backspaceCompare(s string, t string) bool {

	a := backspace(s)
	b := backspace(t)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func backspace(s string) []int {
	stack := NewIntStack(len(s))
	b := '#'
	for _, c := range s {
		if c == b {
			stack.Pop()
		} else {
			stack.Push(int(c))
		}
	}
	// if stack.top == -1 {
	// 	return stack.data[:0]
	// }
	return stack.data[:stack.top+1]
}
