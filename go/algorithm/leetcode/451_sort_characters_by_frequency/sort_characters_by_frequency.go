package leetcode451

import (
	"container/heap"
	"strings"
)

type Node struct {
	Str   string
	Count int
	Index int
}

type priorityQueue []*Node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Count > pq[j].Count
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	x.Index = -1
	*pq = old[0 : n-1]
	return x
}

func (pq *priorityQueue) update(node *Node, value string, priority int) {
	node.Str = value
	node.Count = priority
	heap.Fix(pq, node.Index)
}

func frequencySort(s string) string {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}

	pq := make(priorityQueue, len(m))
	i := 0
	for k, v := range m {
		pq[i] = &Node{k, v, i}
		i++
	}
	heap.Init(&pq)

	var ans string
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		ans += strings.Repeat(node.Str, node.Count)
	}
	return ans

}
