package lru

import "container/list"

/*
   原理：
	添加元素时，放到链表头
	缓存命中，将元素移动到链表头
	缓存满了之后，将链表尾的元素删除
   实现：
	使用双向链表保存数据
   	使用hash实现O(1)的访问
*/

type Cache struct {
	MaxEntries int                           // 数目限制， 0是无限制
	ll         *list.List                    // 使用链表保存数据
	cache      map[interface{}]*list.Element // 使用map查询
}

// Key 是任何可以比较的值
type Key interface{}

// 链表中存储的数据
type entry struct {
	key   Key
	value interface{}
}

// New 创建新的cache对象
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// Add 添加新的值到cache中
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		val.Value.(*entry).value = value
		return
	}

	// 添加数据到链表头部
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		// 满了删除最后访问的元素
		c.RemoveOldest()
	}
}

// Get 从cache中获取数据
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		return val.Value.(*entry).value, true
	}
	return
}

// Remove 删除指定key的元素
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if val, ok := c.cache[key]; ok {
		c.removeElement(val)
	}
}

// RemoveOldest 删除最后访问的元素
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
}

// Len 缓存数
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}
