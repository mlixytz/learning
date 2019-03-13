package leetcode397

/*
   偶数直接右移一位
   奇数 低3位与011按位与，结果为1便-1， 否则+1。目的是保持更少的1位
*/

func integerReplacement(n int) int {
	var count int

	for n != 1 {
		// 偶数
		if n&1 == 0 {
			n >>= 1
		} else if n&3 == 1 || n == 3 {
			n--
		} else {
			n++
		}
		count++
	}
	return count
}
