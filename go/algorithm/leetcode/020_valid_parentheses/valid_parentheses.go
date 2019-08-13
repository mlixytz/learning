package leetcode020

import (
	"container/list"
)

func isValid(s string) bool {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack.PushBack(')')
		case '{':
			stack.PushBack('}')
		case '[':
			stack.PushBack(']')
		default:
			if stack.Len() == 0 || stack.Remove(stack.Back()) != int32(s[i]) {
				return false
			}
		}
	}
	return stack.Len() == 0
}
