package leetcode187

func findRepeatedDnaSequences(s string) []string {
	words := make(map[string]bool)
	repeated := make(map[string]bool)
	ans := make([]string, 0)
	for i := 0; i < len(s)-9; i++ {
		_, ok1 := words[s[i:i+10]]
		_, ok2 := repeated[s[i:i+10]]
		if ok1 {
			if !ok2 {
				repeated[s[i:i+10]] = true
				ans = append(ans, s[i:i+10])
			}
		} else {
			words[s[i:i+10]] = true
		}
	}
	return ans
}
