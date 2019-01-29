package linkedlist

import "container/list"

/*
   双链表排序
*/

// QuickSort 双链表快排
func QuickSort(dlist *list.List) *list.List {
	if dlist == nil {
		return nil
	}
	sort(dlist.Front(), dlist.Back())
	return dlist
}

func sort(p *list.Element, r *list.Element) {
	if p != nil && r != nil && p != r && r != p.Next() {
		pivotIndex := partition(p, r)
		sort(p, pivotIndex.Prev())
		sort(pivotIndex.Next(), r)
	}
}

func partition(p *list.Element, r *list.Element) *list.Element {
	pivot := r.Value.(int)

	i := p.Prev()

	for j := p; j != r; j = j.Next() {
		if j.Value.(int) <= pivot {
			if i == nil {
				i = p
			} else {
				i = i.Next()
			}
			i.Value, j.Value = j.Value, i.Value
		}
	}
	if i == nil {
		i = p
	} else {
		i = i.Next()
	}
	i.Value, r.Value = r.Value, i.Value

	return i
}
