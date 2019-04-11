package leetcode058

func lengthOfLastWord(s string) int {
	ret := 0
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == ' ' {
			if ret > 0 {
				return ret
			}
		} else {
			ret++
		}
	}
	return ret
}
