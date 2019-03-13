package leetcode933

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{list.New()}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue.PushBack(t)
	t -= 3000
	for this.queue.Front().Value.(int) < t {
		this.queue.Remove(this.queue.Front())
	}
	return this.queue.Len()
}
