package leetcode011

func maxArea(height []int) int {
	i, j, ans, area := 0, len(height)-1, 0, 0
	for i < j {
		if height[i] > height[j] {
			area = height[j] * (j - i)
			j--

		} else {
			area = height[i] * (j - i)
			i++
		}
		if area > ans {
			ans = area
		}
	}
	return ans
}
