package leetcode801

func minSwap(A []int, B []int) int {
	size := len(A)
	s1, n1 := 1, 0
	for i := 1; i < size; i++ {
		s2, n2 := 2001, 20001
		if A[i-1] < A[i] && B[i-1] < B[i] {
			s2 = s1 + 1
			n2 = n1
		}
		if A[i-1] < B[i] && B[i-1] < A[i] {
			s2 = min(s2, n1+1)
			n2 = min(n2, s1)
		}
		s1 = s2
		n1 = n2
	}
	return min(n1, s1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
