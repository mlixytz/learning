package leetcode467

func findSubstringInWraproundString(p string) int {
	n := len(p)
	if n == 0 {
		return 0
	}
	count := [26]int{}
	maxLength := 0
	for i := 0; i < n; i++ {
		if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
			maxLength++
		} else {
			maxLength = 1
		}
		if maxLength > count[p[i]-97] {
			count[p[i]-97] = maxLength
		}
	}
	sum := 0
	for i := 0; i < 26; i++ {
		sum += count[i]
	}
	return sum
}
