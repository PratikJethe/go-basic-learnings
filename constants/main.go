package main

import "fmt"

func main() {
	const PublicConstant = "PublicConstant"
	// const privateConstant =  math.Sin(1.75)
	const privateConstant = "privateConstant"
	// privateConstant = "new Value" // not allowed to mutate
	// variable shadwing also works with constant

	const a = 1
	const b int8 = 1

	fmt.Printf("%v,%T", a, a)
	// we could add a(int2) + b(int8) because a is const and compiler allows it
	fmt.Println(a + b)

	const (
		d = iota
		e = iota
		f // compiler will follow pattern and assion iota to f
	)

	fmt.Println(d, e, f)

}
