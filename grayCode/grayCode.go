package graycode

func grayCodeUtil(res *[]int, n int, num []int) {
	if n == 0 {
		*res = append(*res, num[0])
		return
	}

	// ignore the bit.
	grayCodeUtil(res, n-1, num)

	// invert the bit.
	num[0] = num[0] ^ (1 << (n - 1))

	// generate the next Gray code.
	grayCodeUtil(res, n-1, num)
}

func grayCode(n int) []int {
	res := []int{}

	num := []int{0}
	grayCodeUtil(&res, n, num)
	return res
}
