package main

import (
	"fmt"
	"time"
)

func main() {
	c := [4]chan int{}
	for i := 0; i < 4; i++ {
		c[i] = make(chan int)
	}

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				<-c[i]
				fmt.Println(i + 1)
				time.Sleep(time.Second)
				if i+1 == 4 {
					c[i-3] <- 1
				} else {
					c[i+1] <- 1
				}
			}
		}(i)
	}
	c[0] <- 1
	time.Sleep(20 * time.Second)
}
