package leetcode922

func sortArrayByParityII(A []int) []int {
	i, j := 0, 1 // i代表偶数，j代表奇数
	for i <= len(A)-2 && j <= len(A)-1 {
		// 如果偶数索引上不是偶数
		if A[i]%2 != 0 {
			for j <= len(A)-1 {
				if A[j]%2 != 1 {
					A[i], A[j] = A[j], A[i]
					j += 2
					break
				}
				j += 2
			}
		}
		i += 2
	}
	return A
}
