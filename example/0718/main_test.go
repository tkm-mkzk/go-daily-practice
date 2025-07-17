package main

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutinesComplete(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
	// 正常終了
	case <-time.After(time.Second):
		t.Error("Goroutines did not complete in time")
	}
}
