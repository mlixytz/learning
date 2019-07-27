package leetcode739

import "container/list"

func dailyTemperatures(T []int) []int {
	stack := list.New()
	ret := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for stack.Len() > 0 {
			top := stack.Back().Value.(int)
			if T[i] <= T[top] {
				break
			}
			ret[top] = i - top
			stack.Remove(stack.Back())
		}
		stack.PushBack(i)
	}
	return ret
}
