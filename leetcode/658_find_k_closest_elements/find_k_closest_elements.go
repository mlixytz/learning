package leetcode658

func findClosestElements(arr []int, k int, x int) []int {
	size := len(arr)
	l, h := 0, size-1
	for l < h {
		m := l + (h-l)/2
		if arr[m] < x {
			l = m + 1
		} else {
			h = m
		}
	}

	if h < 1 {
		return arr[0:k]
	}

	count, i, j := 0, h-1, h
	for count < k {
		if i < 0 {
			j++
		} else if j > size-1 {
			i--
		} else if x-arr[i] <= arr[j]-x {
			i--
		} else {
			j++
		}
		count++
	}
	return arr[i+1 : j]
}
