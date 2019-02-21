package leetcode461

/*
   位计数算法：Brian Kernighan's Algorithm, O(lg N)。
   原因：每次消除1的最低有效位
*/

func hammingDistance(x int, y int) int {
	result := x ^ y
	ans := 0
	for result != 0 {
		ans++
		result &= result - 1
	}
	return ans
}
