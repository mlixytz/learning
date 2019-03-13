package utils

// IsArrayEqual 判断两个array是否相等
func IsArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IsBytesEqual 判断两个array是否相等
func IsBytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IsArrayNumerEqual 判断两个array中的元素是否相等
func IsArrayNumerEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := 0; i < len(a); i++ {
		m1[a[i]]++
		m2[b[i]]++
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if val, ok := m2[k]; !ok || val != v {
			return false
		}
	}
	return true
}
