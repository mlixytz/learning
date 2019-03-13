package leetcode817

/* 重点：数组转化为map查询。时间复杂度从O(n*m) 变为 O(n+m) */

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	if head == nil || len(G) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, val := range G {
		m[val] = true
	}

	var count int
	var flag bool
	for head != nil {
		_, ok := m[head.Val]
		if ok && !flag {
			flag = true
			count++
		} else if !ok {
			flag = false
		}
		head = head.Next
	}
	return count
}
