package divide_conquer

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	p := 1.0

	for n > 0 {
		if n&1 == 1 {
			p *= x
			// n -= 1
		}
		x *= x
		n >>= 1
	}
	return p
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
