package leetcode389

/*
   相同数字 异或 为0
   0 与任何数字 异或 为数字本身
*/

func findTheDifference(s string, t string) byte {
	var ans byte
	var i int
	for ; i < len(s); i++ {
		ans ^= t[i] ^ s[i]
	}
	ans ^= t[i]
	return ans
}
