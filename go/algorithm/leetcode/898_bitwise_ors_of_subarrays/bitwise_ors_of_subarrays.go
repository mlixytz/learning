package leetcode898

/*
   dp需要去重，否则会超时
*/

func subarrayBitwiseORs(A []int) int {
	res := make(map[int]bool)
	cur := make(map[int]bool)  // 0到前一位元素 所有可能的结果 不会超过
	cur2 := make(map[int]bool) // 0到当前元素 所有可能的结果
	for _, val := range A {
		cur2 = map[int]bool{val: true} // 每轮都包含 自身
		for v := range cur {           // 前一轮 所有结果与当前元素匹配
			cur2[val|v] = true
		}
		cur = cur2 // 得到本轮所有可能的结果
		for v := range cur {
			res[v] = true
		}
	}
	return len(res)
}
