package leetcode692

func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int)
	for _, val := range words {
		counter[val]++
	}
	pql := k
	if len(counter) < k {
		pql = len(counter)
	}
	pq := make(priorityQueue, 0)
	for key, v := range counter {
		pq.insert(element{key, v})
		if len(pq) > pql {
			pq.extractMin()
		}
	}
	ans := make([]string, pql)
	for i := pql - 1; i >= 0; i-- {
		ans[i] = pq.extractMin().key
	}
	return ans
}

type element struct {
	key      string
	priority int
}

type priorityQueue []element

func (pq *priorityQueue) insert(elem element) {
	*pq = append(*pq, elem)
	i := len(*pq) - 1
	p := parent(i)
	for p >= 0 {
		if (*pq)[i].priority < (*pq)[p].priority {
			(*pq)[i], (*pq)[p] = (*pq)[p], (*pq)[i]
			i = p
			p = parent(p)
		} else if (*pq)[i].priority == (*pq)[p].priority && (*pq)[i].key > (*pq)[p].key {
			(*pq)[i], (*pq)[p] = (*pq)[p], (*pq)[i]
			i = p
			p = parent(p)
		} else {
			break
		}
	}
}

func (pq *priorityQueue) extractMin() element {
	n := len(*pq)
	min := (*pq)[0]
	if n > 1 {
		(*pq)[0] = (*pq)[n-1]
		*pq = (*pq)[:n-1]
		minHeapify(*pq, 0)
	}
	return min
}

func parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	}
	return (i - 1) / 2
}

func minHeapify(elems []element, i int) {
	n := len(elems)
	left := 2*i + 1
	right := 2*i + 2
	minIndex := i
	if left < n {
		if elems[minIndex].priority > elems[left].priority {
			minIndex = left
		} else if elems[minIndex].priority == elems[left].priority && elems[minIndex].key < elems[left].key {
			minIndex = left
		}
	}
	if right < n {
		if elems[minIndex].priority > elems[right].priority {
			minIndex = right
		} else if elems[minIndex].priority == elems[right].priority && elems[minIndex].key < elems[right].key {
			minIndex = right
		}
	}
	if minIndex != i {
		elems[i], elems[minIndex] = elems[minIndex], elems[i]
		minHeapify(elems, minIndex)
	}
}
