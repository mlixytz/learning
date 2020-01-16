package leetcode476

/*
   num		00000101
   mask		11111000
   ~mask^num	00000010
*/

func findComplement(num int) int {
	mask := ^0

	for mask&num != 0 {
		mask <<= 1
	}

	return ^mask ^ num
}
