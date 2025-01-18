package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i, j := 0, 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i == 0 {
			time.Sleep(1 * time.Second)
		}
		j = 1
		fmt.Println("Done")
		wg.Done()
	}()
	go func() {
		for j == 0 {
			time.Sleep(1 * time.Second)
		}
		i = 1
		fmt.Println("Done")
		wg.Done()
	}()

	wg.Wait()
}
