package leetcode241

import "strconv"

func diffWaysToCompute(input string) []int {
	ways := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(input[0:i])
			right := diffWaysToCompute(input[i+1:])
			for _, vl := range left {
				for _, vr := range right {
					switch input[i] {
					case '+':
						ways = append(ways, vl+vr)
					case '-':
						ways = append(ways, vl-vr)
					case '*':
						ways = append(ways, vl*vr)
					}
				}
			}
		}
	}
	if len(ways) == 0 {
		temp, _ := strconv.Atoi(input)
		ways = append(ways, temp)
	}
	return ways
}
