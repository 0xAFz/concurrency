package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		fmt.Println("Writer is done!")
		close(ch)
	}()

	go func() {
		time.Sleep(time.Second * 1)
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second * 2)
}
