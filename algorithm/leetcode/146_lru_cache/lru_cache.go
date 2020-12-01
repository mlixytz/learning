package leetcode146

type Node struct {
	key   string
	value int
	left  *Node
	right *Node
}

type LRUCache struct {
	maxSize int
	head    *Node
	tail    *Node
	index   map[string]*Node
}

func Constructor(capacity int) LRUCache {
	if capacity < 1 {
		panic("Capacity is not valid, request greater than 1")
	}
	return LRUCache {
		maxSize: capacity,
		index: make(map[string]*Node)
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this[key]; !ok {
		return
	}
}

func (this *LRUCache) Put(key int, value int) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
