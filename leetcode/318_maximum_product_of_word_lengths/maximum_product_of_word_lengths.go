package leetcode318

func maxProduct(words []string) int {
	binStrs := make([]int, len(words))
	for i, val := range words {
		var tmp int
		for j := 0; j < len(val); j++ {
			tmp |= (1 << (val[j] - 'a'))
		}
		binStrs[i] = tmp
	}

	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if binStrs[i]&binStrs[j] == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}
