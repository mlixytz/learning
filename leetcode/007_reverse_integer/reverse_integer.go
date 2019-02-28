package leetcode007

func reverse(x int) int {
	result := 0
	for x != 0 {
		tail := x % 10
		newResult := result*10 + tail
		if int32(newResult)/10 != int32(result) {
			return 0
		}
		result = newResult
		x /= 10
	}
	return result
}
