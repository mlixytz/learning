package leetcode763

func partitionLabels(S string) []int {
	m := [26]int{} // 记录每个char出现的最后位置
	for i := 0; i < len(S); i++ {
		m[S[i]-'a'] = i
	}
	res := make([]int, 0)
	var start, last int
	for i := 0; i < len(S); i++ {
		if m[S[i]-'a'] > last {
			last = m[S[i]-'a']
		}
		if last == i {
			res = append(res, last-start+1)
			start = last + 1
		}
	}
	return res
}
