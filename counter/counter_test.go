package counter

import (
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	counter := Counter{values: make(map[string]int64)}
	for i := 0; i < 500; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Incr("test")
			}
		}()
	}
	time.Sleep(time.Second * 3)
	if counter.Get("test") != 50000 {
		t.Error("test error")
	}
}
