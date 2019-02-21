package leetcode986

type Interval struct {
	Start int
	End   int
}

func intervalIntersection(A []Interval, B []Interval) []Interval {
	var i, j int
	ans := make([]Interval, 0)
	var a1, a2, b1, b2, s, e int
	for i < len(A) && j < len(B) {
		a1, a2 = A[i].Start, A[i].End
		b1, b2 = B[j].Start, B[j].End
		if a1 < b1 {
			s = b1
		} else {
			s = a1
		}
		if a2 < b2 {
			e = a2
			i++
		} else {
			j++
			e = b2
		}
		if s <= e {
			ans = append(ans, Interval{s, e})
		}
	}
	return ans
}
