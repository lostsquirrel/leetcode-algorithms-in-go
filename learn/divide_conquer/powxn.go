package divide_conquer

func myPow(x float64, n int) float64 {
	p := 1.0
	for i := 0; i < n; {
		p *= x
	}
	return p
}
