package bitmanipulation

/*
   选择恰当的bits
   0xaaaaaaaa		10101010101010101010101010101010
   0x55555555		01010101010101010101010101010101
   0xcccccccc		11001100110011001100110011001100
   0x33333333		00110011001100110011001100110011
   0xf0f0f0f0		11110000111100001111000011110000
   0x0f0f0f0f		00001111000011110000111100001111
*/

// 数字范围按位与
func rangeBitwiseAnd(m, n int) int {
	var i uint32
	for m != n {
		m >>= 1
		n >>= 1
		i++
	}
	return m << i
}

// 二进制中1的个数
func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

// 0x55555555与n按位与 得到奇数位的二进制码
// 4的幂 部分奇数位应该是1
func isPowerOfFour(n int) bool {
	return (n&(n-1)) == 0 && (n&0x55555555) != 0
}
