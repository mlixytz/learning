package leetcode042

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ret, left, right, curmaxleft, curmaxright := 0, 0, len(height)-1, height[0], height[len(height)-1]
	for left < right {
		if height[left] <= height[right] {
			if curmaxleft < height[left] {
				curmaxleft = height[left]
			} else {
				ret += curmaxleft - height[left]
			}
			left++
		} else {
			if curmaxright < height[right] {
				curmaxright = height[right]
			} else {
				ret += curmaxright - height[right]
			}
			right--
		}
	}
	return ret
}
