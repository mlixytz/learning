package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child := context.WithValue(ctx, "hello", "world")

	go func() {
		for {
			select {
			case <-child.Done():
				fmt.Println("It's over")
				return
			default:
				val := child.Value("hello")
				fmt.Println("hello: ", val)
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	time.Sleep(10 * time.Second)
}
