package lru

import (
	"sync"
)

//  实现：
// 	1.使用双向链表存储数据，查询或插入元素后，将元素移动到链表末尾。
//    如果链表长度大于lru最大长度，则从头部删除元素
//  2.使用map作为索引，key保存元素的key，value保存元素的指针

type Node struct {
	key   string
	value int32
	left  *Node
	right *Node
}

type LRU struct {
	mu     sync.Mutex
	size   int32
	maxLen int32
	head   *Node
	tail   *Node
	index  map[string]*Node
}

func (l *LRU) Get(key string) *Node {
	l.mu.Lock()
	defer l.mu.Unlock()
	node, ok := l.index[key]
	if !ok {
		return nil
	}
	// 当只有一个元素或当前节点为尾节点时，直接返回
	if node.right == nil || l.size == 1 {
		return node
	}

	// 当前节点是头节点时
	if node.left == nil {
		node.right.left = nil
		l.head = node.right
	} else {
		node.left.right = node.right
		node.right.left = node.left
	}
	node.left = l.tail
	node.right = nil
	l.tail.right = node
	l.tail = node

	return node
}

func (l *LRU) Put(key string, value int32) {
	l.mu.Lock()
	defer l.mu.Unlock()
	// LRU的index为空则初始化lru
	if l.index == nil {
		l.index = make(map[string]*Node)
	}
	n := &Node{
		key:   key,
		value: value,
	}
	if l.size == 0 {
		l.head = n
		l.tail = n
		l.size++
	} else {
		n.left = l.tail
		l.tail.right = n
		l.tail = n
		if _, ok := l.index[key]; !ok {
			l.size++
			if l.size > l.maxLen {
				delete(l.index, l.head.key)
				l.head = l.head.right
				l.head.left = nil
				l.size--
			}
		}
	}
	l.index[key] = n
}
