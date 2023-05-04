package main

import (
	"fmt"
	// "runtime"
	"sync"
)
// BEST PRACTICES
// Dont create goroutines in libraries let consumer handle concurrecy
// When creating goroutines, know how it will end (saves memory) 

func main() {
	// wait makes program to wait until all goroutines are executed
	// without waitgroup main function will finish before goroutines are exeuted completely
	var wg sync.WaitGroup
	wg.Add(1)     // we add 1 to wg
	go hello(&wg) // running function as gorouitine
	wg.Wait()     // wait until all wg are done

	//synchronization problems with goroutines

	// This code is not thread safe since multiple goroutines access count read it and update it. this will result into inconsistent result without mutex
	var count int = 0
	// runtime.GOMAXPROCS(1)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			count = count + 1
			wg.Done()
		}()

	}
	fmt.Println("Count without mutex ", count)
	var count2 int = 0
	var mu sync.Mutex

	for i := 0; i < 20; i++ {

		go func() {
			wg.Add(1)
			mu.Lock()
			count2 = count2 + 1
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Count with mutex ", count)
	// fmt.Println(runtime.GOMAXPROCS(-1))

}

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hello World")
	wg.Done() // decrementing 1 count
}
