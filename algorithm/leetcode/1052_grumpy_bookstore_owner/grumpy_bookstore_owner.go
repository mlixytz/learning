package leetcode1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		sum += customers[i] * (1 - grumpy[i])
	}
	window := 0
	start := 0
	for start < X {
		window += customers[start] * grumpy[start]
		start++
	}
	maxWindow := window
	for i := start; i < len(customers); i++ {
		window -= customers[i-X] * grumpy[i-X]
		window += customers[i] * grumpy[i]
		if window > maxWindow {
			maxWindow = window
		}
	}
	return sum + maxWindow
}
