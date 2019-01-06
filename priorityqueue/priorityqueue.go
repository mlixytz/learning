package priorityqueue

import (
	"errors"
	"math"
)

// Item 优先级队列中的元素
type Item struct {
	Val      int
	Priority int
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	Items []Item
}

func (pq *PriorityQueue) insert(item Item) {
	tmp := Item{item.Val, math.MinInt64}
	pq.Items = append(pq.Items, tmp)
	pq.increaseKey(len(pq.Items)-1, item.Priority)
}

func (pq *PriorityQueue) maximum() int {
	return pq.Items[0].Val
}

func (pq *PriorityQueue) extractMax() (int, error) {
	if len(pq.Items) < 1 {
		return 0, errors.New("heap underflow")
	}
	max := pq.Items[0]
	pq.Items[0] = pq.Items[len(pq.Items)-1]
	pq.Items = pq.Items[:len(pq.Items)-1]
	maxHeapify(pq.Items, 0)
	return max.Val, nil
}

func (pq *PriorityQueue) increaseKey(index, priority int) error {
	if priority < pq.Items[index].Priority {
		return errors.New("new priority is smaller than current priority")
	}
	pq.Items[index].Priority = priority

	for index > 0 && pq.Items[index].Priority > pq.Items[parent(index)].Priority {
		pq.Items[index], pq.Items[parent(index)] = pq.Items[parent(index)], pq.Items[index]
		index = parent(index)
	}
	return nil
}

func parent(i int) int {
	// 右子叶
	if i%2 == 0 {
		return (i - 2) / 2
	}
	// 左子叶
	return (i - 1) / 2
}

func maxHeapify(items []Item, i int) {
	n := len(items)
	left := i << 2
	right := left + 1
	largest := i
	if left < n && items[left].Priority > items[largest].Priority {
		largest = left
	}
	if right < n && items[right].Priority > items[largest].Priority {
		largest = right
	}
	if largest != i {
		items[i], items[largest] = items[largest], items[i]
		maxHeapify(items, largest)
	}

}
