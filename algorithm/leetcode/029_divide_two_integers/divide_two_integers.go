package leetcode029

import "math"

func divide(dividend int, divisor int) int {
	l, h := 0, abs(dividend)
	for l <= h {
		m := l + (h-l)>>1
		if mul(m, divisor) > abs(dividend) {
			h = m - 1
		} else if mul(m, divisor) < abs(dividend) {
			l = m + 1
		} else {
			return arrangeResult(dividend, divisor, m)
		}
	}
	return arrangeResult(dividend, divisor, h)
}

// best solution
func divideBestSolution(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = false
	}
	dvd := abs(dividend)
	dvs := abs(divisor)
	res := 0
	for dvd >= dvs {
		temp, mul := dvs, 1
		for dvd >= temp<<1 {
			temp <<= 1
			mul <<= 1
		}
		dvd -= temp
		res += mul
	}
	if !sign {
		return -res
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mul(a int, b int) int {
	result := 0
	a = abs(a)
	b = abs(b)
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}

func arrangeResult(a int, b int, c int) int {
	result := c
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		if c > 0 {
			result = 0 - result
		}
	}

	if result < -2147483648 || result > 2147483647 {
		return 2147483647
	}
	return result

}
