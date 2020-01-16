package leetcode744

func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)-1
	for l <= h {
		m := l + (h-l)/2
		if letters[m] < target {
			l = m + 1
		} else if letters[m] > target {
			h = m - 1
		} else {
			for m < len(letters)-1 {
				m++
				if letters[m] != target {
					return letters[m]
				}
			}
			return letters[0]

		}
	}
	if l >= len(letters) {
		return letters[0]
	}
	return letters[l]
}
