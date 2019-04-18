package leetcode049

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, val := range strs {
		ca := []byte(val)
		sort.Slice(ca, func(i, j int) bool { return ca[i] < ca[j] })
		m[string(ca)] = append(m[string(ca)], val)
	}
	ret := make([][]string, 0, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}
