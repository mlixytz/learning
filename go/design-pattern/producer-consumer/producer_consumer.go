package producerconsumer

import "fmt"

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Product: ", i)
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Consume: ", v)
	}
}
