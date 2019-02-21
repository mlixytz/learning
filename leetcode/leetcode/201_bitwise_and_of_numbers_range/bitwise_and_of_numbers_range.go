package leetcode201

/*
   范围区间求余，将start end 全部右移， 直到两个数字相同
   再左移相同的位数，就是最终结果
*/

func rangeBitwiseAnd(m int, n int) int {
	var off uint32
	for m != n {
		m >>= 1
		n >>= 1
		off++
	}
	return m << off
}
