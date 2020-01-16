package leetcode030

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	wl, wsl := len(words[0]), len(words)
	if wl*wsl > len(s) {
		return []int{}
	}
	m := make(map[string]int)
	for _, val := range words {
		m[val]++
	}
	ret := make([]int, 0)
	for i := 0; i <= len(s)-wl*wsl; i++ {
		val, ok := m[s[i:i+wl]]
		if ok && val > 0 {
			tmp := make(map[string]int)
			for k, v := range m {
				tmp[k] = v
			}
			tmp[s[i:i+wl]]--
			isbreak, count := false, 1
			for j := i + wl; (wsl-count > 0) && j <= len(s)-wl*(wsl-count); j += wl {
				_val, _ok := tmp[s[j:j+wl]]
				if _ok && _val > 0 {
					count++
					tmp[s[j:j+wl]]--
				} else {
					isbreak = true
					break
				}
			}
			if !isbreak {
				ret = append(ret, i)
			}
		}
	}
	return ret
}
