package leetcode762

func countPrimeSetBits(L int, R int) int {
	primes := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true}
	var count int
	for i := L; i <= R; i++ {
		if _, ok := primes[hammingWeight(i)]; ok {
			count++
		}
	}
	return count
}

func hammingWeight(n int) int {
	var count int
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
