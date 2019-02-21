package leetcode860

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five < 1 {
				return false
			}
			ten++
			five--
		} else {
			if ten < 1 {
				if five < 3 {
					return false
				}
				five -= 3
			} else if five >= 1 {
				ten--
				five--
			} else {
				return false
			}
		}
	}
	return true
}
