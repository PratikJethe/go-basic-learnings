package main

import "fmt"

func main() {
	//ARRAYS
	grades := [5]int{1, 2}
	// grades := [...]int{1, 2} automatic size
	// grades := [5]int{1, 2,3,4,5,6} overflow not allowed when size provided
	fmt.Println(grades)

	// var flowers [1]string
	// flowers[0] = "rose"
	// fmt.Println(flowers)
	// flowers2 := flowers // arrays are copied by value and not reference

	// flowers2[0] = "lily"
	// fmt.Println(flowers, flowers2)

	//SLICE

	var flowers = []string{"rose"}
	fmt.Println(flowers)
	flowers2 := flowers // slices are copied by reference

	flowers2[0] = "lily"
	fmt.Println(flowers, flowers2) 

	flowers2 =  append(flowers2, "rose")
	fmt.Println(flowers, flowers2) 

	

}
