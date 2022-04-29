package divide_conquer

func myPow(x float64, n int) float64 {
	cache := make(map[int]float64)
	cache[0] = 1.0
	cache[1] = x
	if n < 0 {
		return 1 / pow(cache, x, -n)
	}
	return pow(cache, x, n)
}

func pow(cache map[int]float64, x float64, n int) float64 {
	r, ok := cache[n]
	if ok {
		return r
	}
	h := n / 2
	var x1 float64
	if isEven(n) {
		x1 = pow(cache, x, h) * pow(cache, x, h)
	} else {
		x1 = pow(cache, x, h) * pow(cache, x, h) * x
	}
	cache[n] = x1
	return x1
}

func isEven(n int) bool {
	return n%2 == 0
}
