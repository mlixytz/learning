package leetcode925

func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 {
		return len(typed) == 0
	}
	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if i > 0 && typed[j] != name[i-1] {
			return false
		}
	}
	return i == len(name)
}
