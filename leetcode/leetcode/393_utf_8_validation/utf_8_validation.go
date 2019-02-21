package leetcode393

/* 关键点：穷举列出所有可能的utf-8字节数 */

func validUtf8(data []int) bool {
	var count int
	for _, val := range data {
		if count == 0 {
			switch {
			case val>>3 == 30:
				count = 3
			case val>>4 == 14:
				count = 2
			case val>>5 == 6:
				count = 1
			case val>>7 == 1:
				return false
			}
		} else {
			if val>>6 != 2 {
				return false
			}
			count--
		}
	}
	return count == 0
}
