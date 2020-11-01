package lru

import (
	"testing"
)

func TestLRU(t *testing.T) {
	l := &LRU{maxLen: 5}

	n := l.Get("a")
	if n != nil {
		t.Error("test error")
	}

	l.Put("a", 1)
	n = l.Get("a")
	if n == nil || n.value != 1 {
		t.Error("test error")
	}

	l.Put("b", 2)
	l.Put("c", 3)
	l.Put("d", 4)

	n = l.Get("a")
	if n == nil || n.value != 1 {
		t.Error("test error")
	}

	n = l.Get("c")
	if n == nil || n.value != 3 {
		t.Error("test error")
	}

	l.Put("e", 5)
	l.Put("f", 6)

	n = l.Get("b")
	if n != nil {
		t.Error("test error")
	}
}
