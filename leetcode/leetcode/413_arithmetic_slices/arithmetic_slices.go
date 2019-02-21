package leetcode413

func numberOfArithmeticSlices(A []int) int {
	var curr, sum int
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			curr++
			sum += curr
		} else {
			curr = 0
		}
	}

	return sum
}
