package leetcode056

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if !isEqual(merge([]Interval{Interval{1, 3}, Interval{2, 6}, Interval{8, 10}, Interval{15, 18}}), []Interval{Interval{1, 6}, Interval{8, 10}, Interval{15, 18}}) {
		fmt.Println(merge([]Interval{Interval{1, 3}, Interval{2, 6}, Interval{8, 10}, Interval{15, 18}}))
		t.Error("test error")
	}
	if !isEqual(merge([]Interval{Interval{1, 4}, Interval{4, 5}}), []Interval{Interval{1, 5}}) {
		fmt.Println(merge([]Interval{Interval{1, 4}, Interval{4, 5}}))
		t.Error("test error")
	}
}

func isEqual(a, b []Interval) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].Start != b[i].Start || a[i].End != b[i].End {
			return false
		}
	}
	return true
}
