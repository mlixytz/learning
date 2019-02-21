package leetcode355

type Twitter struct {
	twitters  map[int]map[int]int  // 3个int分别代表 userId, tweetId, timestamp
	followers map[int]map[int]bool // 2个int分别代表 userId(我自己), userId（关注的人）
	time      int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{make(map[int]map[int]int), make(map[int]map[int]bool), 0}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	t, ok := this.twitters[userId]
	if ok {
		this.time++
		t[tweetId] = this.time
	} else {
		this.time++
		tmp := make(map[int]int)
		tmp[tweetId] = this.time
		this.twitters[userId] = tmp
	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	pq := make(priorityQueue, 0)
	// 自己的推文
	for k, v := range this.twitters[userId] {
		pq.insert(Item{k, v})
		if len(pq) > 10 {
			pq.extractMin()
		}
	}

	// 关注者的推文
	for k := range this.followers[userId] {
		for _k, _v := range this.twitters[k] {
			pq.insert(Item{_k, _v})
			if len(pq) > 10 {
				pq.extractMin()
			}
		}
	}
	ans := make([]int, len(pq))
	for i := len(pq) - 1; i >= 0; i-- {
		ans[i] = pq.extractMin().val
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	f, ok := this.followers[followerId]
	if ok {
		f[followeeId] = true
	} else {
		this.followers[followerId] = make(map[int]bool)
		this.followers[followerId][followeeId] = true
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	_, ok := this.followers[followerId]
	if ok {
		delete(this.followers[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type Item struct {
	val      int
	priority int
}

type priorityQueue []Item

func (pq *priorityQueue) insert(item Item) {
	*pq = append(*pq, item)
	i := len(*pq) - 1
	p := parent(i)
	for p >= 0 && (*pq)[i].priority < (*pq)[p].priority {
		(*pq)[i], (*pq)[p] = (*pq)[p], (*pq)[i]
		i = p
		p = parent(p)
	}
	if len(*pq) > 10 {
		pq.extractMin()
	}
}

func (pq *priorityQueue) extractMin() Item {
	n := len(*pq)
	min := (*pq)[0]
	if n > 1 {
		(*pq)[0] = (*pq)[n-1]
		*pq = (*pq)[:n-1]
		minHeapify(*pq, 0)
	}
	return min
}

func minHeapify(items []Item, i int) {
	n := len(items)
	left := 2*i + 1
	right := 2*i + 2
	minIndex := i
	if left < n && items[left].priority < items[minIndex].priority {
		minIndex = left
	}
	if right < n && items[right].priority < items[minIndex].priority {
		minIndex = right
	}
	if minIndex != i {
		items[i], items[minIndex] = items[minIndex], items[i]
		minHeapify(items, minIndex)
	}
}

func parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	}
	return (i - 1) / 2
}
