package bitmanipulation

/* 保持尽可能多的1位 */

// 找到小于等于N的最大2的幂
func largestPower(N int) int {
	// 右边全部换成1
	N |= N >> 1
	N |= N >> 2
	N |= N >> 4
	N |= N >> 8
	N |= N >> 16
	return (N + 1) >> 1
}

// 翻转二进制位
func reverseBits(n uint32) uint32 {
	var mask, ret uint32 = 1 << 31, 0
	for i := 0; i < 32; i++ {
		if n&1 != 0 {
			ret |= mask
		}
		mask >>= 1
		n >>= 1
	}
	return ret
}
