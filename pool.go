package main

import (
	"fmt"
	"sync"
)

func Pool() {

	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	defer myPool.Put(instance)
	myPool.Get()

	// init the pool with one instance
	myPool.Put(myPool.New())

}
