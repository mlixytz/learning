package leetcode032

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, i)
		} else {
			if s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
		}
	}

	if len(stack) == 0 {
		return len(s)
	}
	longest := 0
	a, b := len(s), 0
	for len(stack) > 0 {
		b = stack[len(stack)-1]
		if a-b-1 > longest {
			longest = a - b - 1
		}
		stack = stack[:len(stack)-1]
		a = b
	}
	if a > longest {
		longest = a
	}
	return longest
}
