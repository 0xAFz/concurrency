package main

import (
	"fmt"
	"sync"
)

func main() {
	// channels are FIFO, first in first out
	ch := make(chan string)

	var wg sync.WaitGroup

	producer := func() {
		defer wg.Done()

		x := []string{"Hello, ", "World!"}

		for _, v := range x {
			ch <- v
		}

		close(ch)
	}

	// when use unbuffered channels at least must have one consumer
	consumer := func() {
		defer wg.Done()
		for v := range ch {
			fmt.Print(v)
		}
	}

	wg.Add(2)

	go producer()
	go consumer()

	wg.Wait()
}
