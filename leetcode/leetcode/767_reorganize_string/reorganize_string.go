package leetcode767

import (
	"sort"
)

type tuple struct {
	Val   int
	Count int
}

func reorganizeString(S string) string {
	n := len(S)
	if n < 2 {
		return S
	}
	var charCount [26]int
	for _, c := range S {
		charCount[c-'a']++
	}
	tp := make([]tuple, 0)
	for i := 0; i < 26; i++ {
		if charCount[i] > (n+1)/2 {
			return ""
		}
		tp = append(tp, tuple{i, charCount[i]})
	}
	sort.Slice(tp, func(i, j int) bool { return tp[i].Count > tp[j].Count })
	sortedCount := make([]int, 26)
	for i := range tp {
		if i > 0 {
			sortedCount[tp[i].Val] = tp[i].Count + sortedCount[tp[i-1].Val]
		} else {
			sortedCount[tp[i].Val] = tp[i].Count
		}
	}
	newStr := make([]byte, n)
	for i := 0; i < n; i++ {
		newStr[sortedCount[S[i]-'a']-1] = S[i]
		sortedCount[S[i]-'a']--
	}
	var res string
	for i, j := 0, (n-1)/2+1; i < (n-1)/2+1; i++ {
		res += string(newStr[i])
		if j < n {
			res += string(newStr[j])
		}
		j++
	}
	return res
}
