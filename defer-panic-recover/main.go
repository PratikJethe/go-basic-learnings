package main

import "fmt"

func main() {

	fmt.Println(1)
	//defer is LIFO
	defer fmt.Println(2)
	defer fmt.Println(3)
    fmt.Println("before")
	panicker()
    fmt.Println("after")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		//recover() detects if program has paniced
		if err := recover(); err != nil {
			fmt.Println("Error:",err)
		}
	}()

	panic("Panicked")
	fmt.Println("done panicing")
}
