package leetcode524

import "strings"

func findLongestWord(s string, d []string) string {
	logestWord := ""
	for _, str := range d {
		if len(str) < len(logestWord) {
			continue
		}
		j := 0
		for i := 0; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}
		if j == len(str) {
			if j > len(logestWord) {
				logestWord = str
			} else {
				if strings.Compare(logestWord, str) > 0 {
					logestWord = str
				}
			}
		}
	}
	return logestWord
}
