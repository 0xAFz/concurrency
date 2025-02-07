package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
		ch <- "writer 1"
	}()

	go func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
		ch <- "writer 2"
	}()

	fmt.Println(<-ch)
}
