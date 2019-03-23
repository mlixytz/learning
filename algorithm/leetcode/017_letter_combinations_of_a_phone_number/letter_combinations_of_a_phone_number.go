package leetcode017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[int][]byte{2: {'a', 'b', 'c'}, 3: {'d', 'e', 'f'}, 4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'}, 6: {'m', 'n', 'o'}, 7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'}, 9: {'w', 'x', 'y', 'z'}}
	ret := make([]string, 0)
	tmpList := make([]byte, 0)
	backtrack(&tmpList, digits, 0, m, &ret)
	return ret
}

func backtrack(tmp *[]byte, digits string, i int, m map[int][]byte, ret *[]string) {
	if i == len(digits) {
		_tmp := make([]byte, len(digits))
		copy(_tmp, *tmp)
		*ret = append(*ret, string(_tmp))
		return
	}
	for _, val := range m[int(digits[i]-'0')] {
		*tmp = append(*tmp, val)
		backtrack(tmp, digits, i+1, m, ret)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
