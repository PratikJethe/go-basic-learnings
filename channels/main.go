package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// ch := make(chan int)
	// // ch := make(chan int,50) //buffered channel
	// wg.Add(2)

	// go func(ch <-chan int) { // making ch read only
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) { // making ch write only
	// 	ch <- 20
	// 	// ch <- 10 // pushing value to channel after first goroutine exits
	// 	wg.Done()
	// }(ch)
	// wg.Wait()
	// // we can close channel by close(ch), pushing values to closed channel will cause error

	//select statement
	logCh := make(chan string)
	closeLog := make(chan bool)

	go func(logch chan<- string, closeLog chan<- bool) {
		logch <- "msg 1"
		logch <- "msg 2"
		logch <- "msg 3"
		closeLog <- true

	}(logCh, closeLog)

	func() {
		for {
			select {
			case msg := <-logCh:
				fmt.Println(msg)

			case <-closeLog:
				fmt.Println("exiting logger")
				return;

			}
		}

	}()

}
