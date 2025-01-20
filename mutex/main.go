package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	x := 0
	mux := sync.RWMutex{}

	read := func() {
		mux.RLock()
		time.Sleep(100 * time.Millisecond)
		mux.RUnlock()
		wg.Done()
	}

	write := func() {
		start := time.Now()
		for !mux.TryLock() {
			time.Sleep(10 * time.Millisecond)
		}

		x++
		mux.Unlock()
		fmt.Printf("wait time: %vms\n", time.Since(start).Milliseconds())
		wg.Done()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
}

//
// func mutex() {
//     wg := sync.WaitGroup{}
//
//     x := 0
//     mux := sync.Mutex{}
//
//     read := func() {
//         mux.Lock()
//         time.Sleep(100 * time.Millisecond)
//         mux.Unlock()
//         wg.Done()
//     }
//
//     write := func() {
//         start := time.Now()
//         for !mux.TryLock() {
//             time.Sleep(10 * time.Millisecond)
//         }
//
//         x++
//         mux.Unlock()
//         fmt.Printf("wait time: %vms\n", time.Since(start).Milliseconds())
//         wg.Done()
//     }
//
//     for i := 0; i < 20; i++ {
//         wg.Add(1)
//         go read()
//     }
//     for i := 0; i < 3; i++ {
//         wg.Add(1)
//         go write()
//     }
//
//     wg.Wait()
// }
//
