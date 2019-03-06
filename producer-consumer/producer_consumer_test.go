package producerconsumer

import (
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	ch1 := make(chan int)
	go producer(ch1)
	go consumer(ch1)
	time.Sleep(time.Second)

	ch10 := make(chan int, 10)
	go producer(ch10)
	go consumer(ch10)
	time.Sleep(time.Second)
}
