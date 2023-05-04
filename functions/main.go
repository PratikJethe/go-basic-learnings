package main

import (
	"fmt"
)

type WalterWhite struct {
	name string
}

func (w WalterWhite) sayMyName() {
	fmt.Println(w.name)

	w.name = "changed name" // wont chnage actual name in struct since struct was passed by value

}
func (w *WalterWhite) sayMyNameReference() {
	fmt.Println(w.name)

	w.name = "changed name" // will chnage actual name in struct since struct was passed by reerence

}

func main() {
	sayMyName("Heisenberg")
	printMultipleNumbers(1, 2, 3, 4, 5, 6, 7)
	returnSum(1, 1)

	value, err := returnDivide(10, 0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(value)

	func() {
		fmt.Println("self executing anonymous function")
	}() //anonymous functiom

	// f() variable func cannot be called before initialization
	var f func() = func() {

		fmt.Println("func as variable")
	}

	f()

	//Struct Methods

	mrWhite := WalterWhite{
		name: "Heisenberg",
	}

	mrWhite.sayMyName()
	fmt.Println(mrWhite.name)

	//BY REFERNCE
	mrWhite2 := WalterWhite{
		name: "Heisenberg",
	}

	mrWhite2.sayMyNameReference()
	fmt.Println(mrWhite2.name)

}

func sayMyName(name string) {
	fmt.Println(name)
}
func printMultipleNumbers(values ...int) {
	fmt.Println(values)
}

func returnSum(a, b int) int {
	return a + b
}

func returnMul(a, b int) (result int) { //result is auto declared by compiler here
	result = a * b
	return
}

func returnDivide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
