package leetcode014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	result := make([]byte, 0)
	for i := 0; i < len(strs[0]); i++ {
		b := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != b {
				return string(result)
			}
		}
		result = append(result, b)
	}
	return string(result)
}
