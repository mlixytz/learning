package leetcode012

func intToRoman(num int) string {
	var ret string
	src := []string{"I", "V", "X", "L", "C", "D", "M"}
	off := 0
	for num > 0 {
		n := num % 10
		if n < 4 {
			c := n
			for c > 0 {
				ret = src[off] + ret
				c--
			}
		} else if n == 4 {
			ret = src[off] + src[off+1] + ret
		} else if n < 9 {
			c := n
			for c-5 > 0 {
				ret = src[off] + ret
				c--
			}
			ret = src[off+1] + ret
		} else {
			ret = src[off] + src[off+2] + ret
		}
		num /= 10
		off += 2
	}
	return ret
}
