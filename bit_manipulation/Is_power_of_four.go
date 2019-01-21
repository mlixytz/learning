package bitmanipulation

/* 0x55555555与n按位与 得到奇数位的二进制码
   4的幂 部分奇数位应该是1 */

func isPowerOfFour(n int) bool {
	return (n&(n-1)) == 0 && (n&0x55555555) != 0
}
