package leetcode135

func candy(ratings []int) int {
	length := len(ratings)
	candys := make([]int, length)
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candys[i] <= candys[i+1] {
			candys[i] = candys[i+1] + 1
		}
	}
	sum := 0
	for i := 0; i < length; i++ {
		sum += candys[i] + 1
	}
	return sum
}
