package priorityqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := new(PriorityQueue)
	i1 := Item{10, 2}
	pq.insert(i1)
	if pq.maximum() != 10 {
		t.Error("test error")
	}
	i2 := Item{11, 3}
	pq.insert(i2)
	if pq.maximum() != 11 {
		t.Error("test error")
	}
	i3 := Item{12, 1}
	pq.insert(i3)
	if pq.maximum() != 11 {
		t.Error("test error")
	}
	pq.increaseKey(2, 5)
	if pq.maximum() != 12 {
		t.Error("test error")
	}
	i4 := Item{100, 4}
	pq.insert(i4)
	if pq.maximum() != 12 {
		t.Error("test error")
	}
	val, err := pq.extractMax()
	if err != nil {
		t.Error("test error")
	} else {
		if val != 12 {
			t.Error("test error")
		}
	}
}
