package leetcode089

/*
	n=3的时候基于n=2，依次类推
	00,01,11,10 -> (000,001,011,010) (110,111,101,100)
	左括号与右括号对称，都是最高位不同
*/

func grayCode(n int) []int {
	res := make([]int, 0)
	res = append(res, 0)
	for i := 0; i < n; i++ {
		k := len(res) - 1
		for j := k; j >= 0; j-- {
			res = append(res, res[j]|1<<uint(i))
		}
	}
	return res
}
