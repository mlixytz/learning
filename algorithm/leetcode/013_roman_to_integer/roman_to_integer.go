package leetcode013

func romanToInt(s string) int {
	ret := 0
	m1 := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	m2 := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	i := 0
	for i < len(s) {
		if i < len(s)-1 && (s[i] == 'I' || s[i] == 'X' || s[i] == 'C') {
			val, ok := m2[s[i:i+2]]
			if ok {
				ret += val
				i += 2
			} else {
				ret += m1[s[i]]
				i++
			}
		} else {
			ret += m1[s[i]]
			i++
		}
	}
	return ret
}
