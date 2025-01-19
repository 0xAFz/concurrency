package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(100)

	// var mux sync.Mutex

	x := 0

	for i := 1; i < 101; i++ {
		runtime.Gosched() // Gosched: Golang scheduler literally telling: "What the fuck is race condition?"
		go func(j int) {
			defer wg.Done()
			// mux.Lock()
			x += j
			// mux.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(x)
}
