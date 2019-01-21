package bitmanipulation

/*
   1.使用^删除完全相同的数字（相同数异或为0）
   2.保留不同位 移除相同位
   3.0与任何数^ 都为数本身
*/

// 使用位计算使两个数相加
func sumOfTwoIntegers(a, b int) int {
	if b == 0 {
		return a
	}
	return sumOfTwoIntegers(a^b, (a&b)<<1)
}

// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，
// 找出 0 .. n 中没有出现在序列中的那个数。 如[0,1,3] 返回2
func missingNumber(nums []int) int {
	/*
		原数组：0 1 2 4
		原索引：0 1 2 3

		因为 a^b^b = a
		所以 0 ^ 0 ^ 1 ^ 1 ^ 2 ^ 2 ^ 3 ^ 4 ^ 4(数组长度) = 3
	*/

	var ret int
	for i, val := range nums {
		ret ^= i ^ val
	}

	ret ^= len(nums)
	return ret
}
