package leetcode725

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	// 求出链表长度
	var length int
	fast := root
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		length += 2
	}
	if fast != nil {
		length++
	}

	p := length / k
	r := length % k
	nh := root
	ans := make([]*ListNode, 0)
	for i := 0; i < k; i++ {
		s := p - 1
		if r > 0 {
			s++
			r--
		}
		for nh != nil && s > 0 {
			nh = nh.Next
			s--
		}
		if root != nil {
			tmp := nh.Next
			nh.Next = nil
			ans = append(ans, root)
			root = tmp
			nh = root
		} else {
			ans = append(ans, root)
		}
	}

	return ans
}
