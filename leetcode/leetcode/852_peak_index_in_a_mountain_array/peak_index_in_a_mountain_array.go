package leetcode852

func peakIndexInMountainArray(A []int) int {
	l, h := 0, len(A)-1
	for l <= h {
		m := l + (h-l)/2
		if m == 0 {
			return 0
		}
		if m == len(A)-1 {
			return m
		}
		if A[m-1] > A[m] && A[m] > A[m+1] {
			h = m - 1
		} else if A[m-1] < A[m] && A[m] < A[m+1] {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}
