package leetcode1053

import "fmt"

func prevPermOpt1(A []int) []int {
	if len(A) < 2 {
		return A
	}
	left, right := len(A)-2, len(A)-1
	for left >= 0 && A[left] <= A[left+1] {
		left--
	}
	if left < 0 {
		return A
	}
	fmt.Println(left, right)

	for A[right] >= A[left] {
		right--
	}
	for A[right] == A[right-1] {
		right--
	}
	fmt.Println(left, right)

	A[left], A[right] = A[right], A[left]
	return A
}
