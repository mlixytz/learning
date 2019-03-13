package lru

import "testing"

func TestSolution(t *testing.T) {
	cache := New(5)
	cache.Add("a1", 10)
	cache.Add("a2", 20)
	cache.Add("a3", 30)
	cache.Add("a4", 40)
	cache.Add("a5", 50)
	cache.Add("a6", 60)
	_, exist := cache.Get("a1")
	if exist {
		t.Error("test error")
	}
	val, exist := cache.Get("a2")
	if !exist {
		t.Error("test error")
	} else {
		if val != 20 {
			t.Error("test error")
		}
	}
	cache.Remove("a2")
	_, exist = cache.Get("a2")
	if exist {
		t.Error("test error")
	}
}
