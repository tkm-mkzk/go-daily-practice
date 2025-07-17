package main

import (
	"fmt"
	"sync"
)

func printNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine", id, "prints number", i)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}
	wg.Wait()
}
