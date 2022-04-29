package divide_conquer

func myPow(x float64, n int) float64 {
	cache := make(map[int]float64)
	cache[0] = 1.0
	cache[1] = x
	return pow(cache, x, n)
}

func pow(cache map[int]float64, x float64, n int) float64 {
	r, ok := cache[n]
	if ok {
		return r
	}
	h := n / 2
	x1 := pow(cache, x, h) * pow(cache, x, n-h)
	cache[n] = x1
	return x1
}
