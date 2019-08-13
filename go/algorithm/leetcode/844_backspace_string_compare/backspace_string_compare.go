package leetcode844

func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	for {
		for back := 0; i >= 0 && (back > 0 || S[i] == '#'); i-- {
			if S[i] == '#' {
				back++
			} else {
				back--
			}
		}

		for back := 0; j >= 0 && (back > 0 || T[j] == '#'); j-- {
			if T[j] == '#' {
				back++
			} else {
				back--
			}
		}

		if i >= 0 && j >= 0 && S[i] == T[j] {
			i--
			j--
		} else {
			return i == -1 && j == -1
		}
	}
}
