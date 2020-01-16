package leetcode136

/* 相同数异或结果位 0
   0 与任何数异或结果位数本身
*/

func singleNumber(nums []int) int {
	ans := 0
	for _, val := range nums {
		ans ^= val
	}
	return ans
}
