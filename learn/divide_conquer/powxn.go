package divide_conquer

func myPow(x float64, n int) float64 {
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	h := n / 2
	return pow(x, h) * pow(x, n-h)
}
