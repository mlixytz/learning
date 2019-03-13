package leetcode003

// lengthOfLongestSubstring my solution
func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	byteMap := make(map[byte]int, 0)

	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		if val, ok := byteMap[s[j]]; ok {
			for k := i; k < val+1; k++ {
				delete(byteMap, s[k])
			}
			i = val + 1
		} else {
			if maxLength < j-i+1 {
				maxLength = j - i + 1
			}
		}
		byteMap[s[j]] = j
	}

	return maxLength
}

// lengthOfLongestSubstringBestSolution best solution
func lengthOfLongestSubstringBestSolution(s string) int {
	maxLength, left := 0, 0

	var letters [256]int
	for i := range letters {
		letters[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if letters[s[i]] >= left {
			left = letters[s[i]] + 1
		} else if i-left+1 > maxLength {
			maxLength = i - left + 1
		}
		letters[s[i]] = i
	}

	return maxLength
}
