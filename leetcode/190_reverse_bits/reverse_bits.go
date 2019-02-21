package leetcode190

func reverseBits(num uint32) uint32 {
	var mask, ans uint32 = 1 << 31, 0
	for i := 0; i < 32; i++ {
		if num&1 != 0 {
			ans |= mask
		}
		mask >>= 1
		num >>= 1
	}
	return ans
}
