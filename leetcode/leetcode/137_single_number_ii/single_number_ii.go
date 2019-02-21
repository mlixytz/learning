package leetcode137

/*
   思路：为每一位二进制进行计数,计数三次归零,那么最后没有归零的数就是出现一次的那个数
   数字第一次出现 添加到ones中，第二次出现，从ones中移除，添加到twos中，第三次出现，从twos中移除
*/

func singleNumber(nums []int) int {
	var ones, twos int
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
