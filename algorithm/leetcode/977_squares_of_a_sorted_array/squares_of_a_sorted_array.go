package leetcode977

func sortedSquares(A []int) []int {
	l, r := 0, len(A)-1
	res := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		leftVal := A[l] * A[l]
		rightVal := A[r] * A[r]
		if leftVal > rightVal {
			res[i] = leftVal
			l++
		} else {
			res[i] = rightVal
			r--
		}
	}
	return res
}
