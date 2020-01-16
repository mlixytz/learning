package leetcode028

func strStr(haystack string, needle string) int {
	var i, j int
	for ; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] {
			j++
		} else if j > 0 {
			i -= j
			j = 0
		}
	}
	if j < len(needle) {
		return -1
	}
	return i - j // for执行完循环体，会接着执行“循环变量迭代”，然后再判断“循环条件”
}
