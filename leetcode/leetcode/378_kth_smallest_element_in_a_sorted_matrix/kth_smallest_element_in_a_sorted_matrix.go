package leetcode378

func kthSmallest(matrix [][]int, k int) int {
	l := matrix[0][0]
	h := matrix[len(matrix)-1][len(matrix)-1]

	for l < h {
		m := l + (h-l)/2
		count := 0
		j := len(matrix) - 1
		for i := 0; i < len(matrix); i++ {
			for j >= 0 && matrix[i][j] > m {
				j--
			}
			count += j + 1
		}
		if count < k {
			l = m + 1
		} else {
			h = m
		}
	}
	return l
}
