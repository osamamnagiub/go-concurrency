package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func Channels() {

	// var dataStream chan interface{}
	// dataStream = make(chan interface{})

	// var recieveOnlyDataStream <-chan interface{}
	// recieveOnlyDataStream = make(<-chan interface{})

	// var sentOnlyDataStream chan<- interface{}
	// sentOnlyDataStream = make(chan<- interface{})

	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}

func Channels2() {

	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels!"
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

func ClosedChannel() {

	stringStream := make(chan string)
	close(stringStream)

	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

func RangeChannels() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}

}

func UnBlockingAllRoutines() {
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func BufferedChannels() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func ChannelOwnership() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func Select() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func SelectWithTimout() {
	c := make(chan interface{})
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}

func SelectDefault() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

// This allows you to exit a select block without blocking. Usually you'll see a default clause used in
// conjunction with a for-select loop. This allows a goroutine to make progress on work
// while waiting for another goroutine to report a result.
func ForSelectLoop() {
	done := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0

Loop:
	for {
		select {
		case <-done:
			break Loop
		default:
		}

		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
