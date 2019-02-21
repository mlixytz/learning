package leetcode405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var ans string
	char := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	n := uint32(num) // 负数转换为正数计算
	for n > 0 {
		ans = char[n&0xf] + ans
		n >>= 4
	}
	return ans
}
