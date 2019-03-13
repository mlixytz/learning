package leetcode008

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	sign, i, base := 1, 0, 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i >= len(str) {
		return 0
	}
	if str[i] == '-' || str[i] == '+' {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		if base > (1<<31-1)/10 || (base == (1<<31-1)/10 && str[i]-'0' > 7) {
			if sign == -1 {
				return -1 << 31
			}
			return 1<<31 - 1
		}
		base = 10*base + int(str[i]-'0')
		i++
	}
	return base * sign
}
