package leetcode038

import "fmt"

func countAndSay(n int) string {
	ret := "1"
	for i := 2; i <= n; i++ {
		count, cur, _ret := 1, ret[0], ""
		for j := 1; j < len(ret); j++ {
			if cur == ret[j] {
				count++
			} else {
				_ret += fmt.Sprintf("%d%c", count, cur)
				count = 1
				cur = ret[j]
			}
		}
		_ret += fmt.Sprintf("%d%c", count, cur)
		ret = _ret
	}
	return ret
}
