package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Second)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2st goroutine sleeping...")
		time.Sleep(2 * time.Second)

	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")

}

func WaitGroup2() {
	var wg sync.WaitGroup
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const num = 5
	wg.Add(5)

	for i := 0; i < num; i++ {
		go hello(&wg, i)
	}

	wg.Wait()
}
