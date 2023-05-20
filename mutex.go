package main

import (
	"fmt"
	"sync"
)

func Mutex() {
	var count int
	var lock sync.Mutex

	increament := func() {
		lock.Lock()
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decreament := func() {
		lock.Lock()
		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increament()
		}()
	}

	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decreament()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

func RWMutex() {
	var count int
	var lock sync.RWMutex

	increament := func() {
		lock.Lock()
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decreament := func() {
		lock.Lock()
		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	reading := func() {
		lock.RLock()
		defer lock.RUnlock()

		fmt.Printf("Reading Count: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increament()
		}()
	}

	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decreament()
		}()
	}

	// Only read
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			reading()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
