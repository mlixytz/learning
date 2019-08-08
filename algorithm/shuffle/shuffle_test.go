package shuffle

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Error(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}))
	t.Error(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}))
	t.Error(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}))
	t.Error(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}))
	t.Error(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}))
}
