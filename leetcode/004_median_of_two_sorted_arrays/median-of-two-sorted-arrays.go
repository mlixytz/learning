package leetcode004

/*

划分数组x，y。划分点为px，py。使得
条件：x[px-1] <= y[py] y[py-1] <= x[px]
公式：(len(x)+len(y)+1)/2 - px = py

使用二分法查找x数组分割点：
if x[midx] > y[py]:
	midx 左移
else:
	midx 右移
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 防止公式中py出现负数
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		px := low + (high-low)/2
		py := (x+y+1)/2 - px

		var maxLeftX, maxLeftY, minRightX, minRightY int
		if px == 0 {
			maxLeftX = -1 << 31
		} else {
			maxLeftX = nums1[px-1]
		}
		if px == x {
			minRightX = 1 << 31
		} else {
			minRightX = nums1[px]
		}

		if py == 0 {
			maxLeftY = -1 << 31
		} else {
			maxLeftY = nums2[py-1]
		}
		if py == y {
			minRightY = 1 << 31
		} else {
			minRightY = nums2[py]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			}
			return float64(max(maxLeftX, maxLeftY))
		} else if maxLeftX > minRightY {
			high = px - 1
		} else {
			low = px + 1
		}
	}
	return 0.0

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
