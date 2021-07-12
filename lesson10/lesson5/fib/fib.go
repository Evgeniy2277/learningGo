package fib

func FibRec(n int) int {
	if n <= 1 {
		return n
	}
	return FibRec(n - 1) + FibRec(n -2)
}
var cache = map[int]int{}
func FibCache(n int) int {
	v, ok := cache[n]
	if !ok {
		v = FibRecCache(n)
		cache[n] = v
	}
	return v
}

func FibRecCache(n int) int {
	if n <= 1 {
		return n
	}
	return FibCache(n -1) + FibCache(n - 2)
}
