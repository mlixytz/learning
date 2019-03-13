package leetcode275

func hIndex(citations []int) int {
	size := len(citations)
	l, h := 0, size-1
	for l <= h {
		m := l + (h-l)/2
		if citations[m] < size-m {
			l = m + 1
		} else if citations[m] > size-m {
			if m < 1 || citations[m-1] <= size-m {
				return size - m
			}
			h = m - 1
		} else {
			return size - m
		}
	}
	return 0
}
