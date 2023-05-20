package main

import (
	"fmt"
	"sync"
)

func main() {
	// go sayHello()

	// go func() {
	// 	println("Hello")
	// }()

	// hello := func() {
	// 	println("Hello")
	// }

	// go hello()

	// time.Sleep(2000)

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func sayHello() {
	println("Hello")
}
