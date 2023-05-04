package main

import (
	"fmt"
	"strconv"
)

// decalred in global scope
var l int = 10
var j int = 22 // decalre j in  global scope


// grouped with var
var (
	name   string  = "name"
	age    int     = 20
	weight float32 = 22.2
)

func main() {

	fmt.Println("j befor initializing in local scope: ", j)

	//ways to decalre variable
	var i int      //not initialized
	var j int = 10 //initialized
	k := 20        //initialized

	var privateVariable string = "private" // variable with first letter lowercase are private to package
	var PublicVariable string = "Public"   // variable with first letter uppercase are public and can be imported
	fmt.Println("j after initializing in local scope: ", j)
	fmt.Println("this is called shadowing, local variable is considered first")

	fmt.Println(i, k,privateVariable,PublicVariable)


	//converting datatypes

	var numFloat float32= 10.2
	var numInt int = int(numFloat)
	var numString string =  strconv.Itoa(numInt)
	fmt.Println(numFloat,numInt,numString)

}
