package general

// return nth fib number
func GetFibNumber(n int) int {
	if n <= 1 {
		return n
	}
	return GetFibNumber(n-1) + GetFibNumber(n-2)
}

func GetFibIter(n int) int {
	f, s := 0, 1
	for i := 0; i < n; i++ {
		f, s = s, f+s
	}
	return f
}
