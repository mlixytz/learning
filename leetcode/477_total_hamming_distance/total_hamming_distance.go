package leetcode477

/*
   举例：（n代表1的数量，m代表0的数量）
   4			0     1     0     0
   14			0     1     1     1
   2			0     0     1     0
   汉明总距离(n*m)       0  +  2  +  2  +  2 = 6
*/

func totalHammingDistance(nums []int) int {
	n := len(nums)
	var i uint
	ans := 0
	for ; i < 32; i++ {
		onebit := 0
		for _, val := range nums {
			onebit += (val >> i) & 1
		}
		ans += onebit * (n - onebit)
	}
	return ans
}
