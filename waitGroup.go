package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGrop() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)

	}()

	wg.Add(1)
	go func() {
		fmt.Println("2st goroutine sleeping...")
		time.Sleep(2)

	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")

}
