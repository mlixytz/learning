package leetcode260

/*
   命出现一次的数字为a, b
   循环第一遍：求出 c = a^b
   找到a与b不同的最后一位: aXorb & (-aXorb)
   然后将这个数组按照最后不同位分成两组：groupA包含a 和groupB不包含a的
   a与groupA求异或得到a， b与groupB求异或得到b
*/

func singleNumber(nums []int) []int {
	aXorb := 0
	for _, val := range nums {
		aXorb ^= val
	}
	lastBit := aXorb & (-aXorb)
	var a, b int
	for _, val := range nums {
		if val&lastBit != 0 {
			a ^= val
		} else {
			b ^= val
		}
	}
	return []int{a, b}
}
