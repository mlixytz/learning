package leetcode002

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	l1 := []int{1, 7, 3}
	l2 := []int{5, 6, 4}
	ln1 := listToNode(l1)
	ln2 := listToNode(l2)
	result := addTwoNumbers(ln1, ln2)
	t.Log(result.Show())
	if result.Show() != "638" {
		t.Error("test Error")
	}
}

// Show ListNode value
func (ln *ListNode) Show() string {
	str := ""
	for {
		if ln.Next != nil {
			str += strconv.Itoa(ln.Val)
			ln = ln.Next
		} else {
			str += strconv.Itoa(ln.Val)
			break
		}
	}
	return str
}

func listToNode(li []int) *ListNode {
	var node *ListNode
	liLen := len(li)
	for i := 0; i < liLen; i++ {
		node = &ListNode{li[liLen-1-i], node}
	}
	return node
}
