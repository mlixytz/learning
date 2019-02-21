package leetcode875

func minEatingSpeed(piles []int, H int) int {

	l, h := 1, maxValInSlice(piles)
	for l < h {
		m := l + (h-l)/2

		times := 0
		for _, val := range piles {
			if val%m == 0 {
				times += val / m
			} else {
				times += val/m + 1
			}
		}
		if times > H {
			l = m + 1
		} else {
			h = m
		}
	}
	return l

}

func maxValInSlice(s []int) int {
	a := -1
	for _, val := range s {
		a = max(val, a)
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
