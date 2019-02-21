package leetcode931

func minFallingPathSum(A [][]int) int {
	size := len(A)
	if size == 1 {
		return A[0][0]
	}
	for i := 1; i < size; i++ {
		for j := 0; j < size; j++ {
			v := A[i-1][j]
			if j > 0 {
				v = min(v, A[i-1][j-1])
			}
			if j < size-1 {
				v = min(v, A[i-1][j+1])
			}
			A[i][j] += v
		}
	}
	return min(A[size-1]...)
}

func min(arr ...int) int {
	v := arr[0]
	for _, j := range arr {
		if j < v {
			v = j
		}
	}
	return v
}
