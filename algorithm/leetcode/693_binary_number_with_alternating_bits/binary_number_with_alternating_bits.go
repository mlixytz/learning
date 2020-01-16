package leetcode693

/*
   n		00001010
   n>>1         00000101
   n = n ^ n>>1	00001111
   n+1		00010000
   n + 1 & n    00000000
*/

func hasAlternatingBits(n int) bool {
	n = n ^ n>>1
	return (n+1)&n == 0
}
