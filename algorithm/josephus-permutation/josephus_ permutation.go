package josephuspermutation

/*
	josephus permutation introduction:
	People are standing in a circle waiting to be executed. Counting begins at a
	specified point in the circle and proceeds around the circle in a specified
	direction. After a specified number of people are skipped, the next person is
	executed. the procedure is repeated with the remaining people, starting with
	the next person, going in the samedirection and skipping the same number of
	people, until only one person remains, and is freed.
*/

type listNode struct {
	Val  int
	Next *listNode
}

// josephusPermutation1 solves the problem with circular linked list
// n represents the number of people, s represents the start position
// of counting, k represents skipped number. It's time complexity is O(n * k)
func josephusPermutation1(n, s, k int) int {
	if n == 1 {
		return 1
	}
	// generate cirular linked list and find the start node
	head := &listNode{Val: 1}
	current, start := head, head
	for i := 0; i < n-1; i++ {
		current.Next = &listNode{Val: i + 2}
		current = current.Next
		s--
		if s == 1 {
			start = current
		}
	}
	current.Next = head
	head = start

	for head != head.Next {
		for count := k; count > 0; count-- {
			head = head.Next
		}
		head.Val = head.Next.Val
		head.Next = head.Next.Next
	}
	return head.Val
}

// josephusPermutation2 solves the problem with recursive traversal
// n represents the number of people, k represents skipped number.
// It's time complexity is O(n)
// f(n, k) = (f(n-1, k) + k) % n + 1
func josephusPermutation2(n, k int) int {
	if n == 1 {
		return 1
	}
	return (josephusPermutation2(n-1, k)+k)%n + 1
}
