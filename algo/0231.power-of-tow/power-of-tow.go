package problem0231

func isPowerOfTwo(n int) bool {
	return n > 0 && ((n-1)&n == 0)
}
