package leetcode242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b := make([]int, 26)
	for i := 0; i < len(s); i++ {
		b[s[i]-'a']++
		b[t[i]-'a']--
	}
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
