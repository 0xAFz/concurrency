package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	c := 0

	taskOne := func() {
		if c == 0 {
			defer func() {
				c = 0
			}()

			c = 1

			fmt.Println("Task One")

			// keep resource busy
			time.Sleep(100 * time.Millisecond)
		}
	}

	taskTwo := func() {
		if c == 0 {
			defer func() {
				c = 0
			}()

			c = 1

			fmt.Println("Task Two")
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	go func() {
		for {
			taskOne()
		}
	}()

	go func() {
		time.Sleep(1 * time.Millisecond)
		for {
			taskTwo()
		}
	}()

	wg.Wait()
}
