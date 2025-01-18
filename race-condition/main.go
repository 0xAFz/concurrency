package main

import "fmt"

func main() {
	for {
		i := 0

		go func() {
			fmt.Println(i)
		}()

		go func() {
			i++
		}()
	}
}
