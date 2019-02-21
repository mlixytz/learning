package leetcode659

func isPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	// tmp 索引对应元素可以被添加到 多少个连续序列尾部
	counter, tmp := make(map[int]int), make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}
	for _, v := range nums {
		if counter[v] == 0 {
			continue
		}
		if tmp[v] > 0 { // 如果该元素可以添加到任意连续序列的尾部
			tmp[v]--
			counter[v]--
			tmp[v+1]++
		} else if counter[v+1] > 0 && counter[v+2] > 0 { // 如果该元素不能添加到任意连续序列的尾部，检查是否可以成为新序列的首元素
			counter[v]--
			counter[v+1]--
			counter[v+2]--
			tmp[v+3]++
		} else {
			return false
		}
	}
	return true
}
