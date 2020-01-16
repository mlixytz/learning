package counter

import "sync"

// Counter 线程安全的计数器
type Counter struct {
	mu     sync.Mutex
	values map[string]int64
}

// Get 获取计数器中的值
func (c *Counter) Get(key string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[key]
}

// Incr 增加计数器中的值
func (c *Counter) Incr(key string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key]++
	return c.values[key]
}
