package bitmanipulation

/*
   计算数字对应二进制中有多少个1
*/

func countOne(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
