package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	go countfunc()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("1st goroutine awake")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("2nd goroutine awake")
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")

}

func countfunc() {
	for {
		fmt.Println(runtime.NumGoroutine())
		time.Sleep(100 * time.Millisecond)
	}
}

