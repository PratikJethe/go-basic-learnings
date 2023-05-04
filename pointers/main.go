package main

import "fmt"

func main() {
 var a int = 42
var b int = a
var aPointer *int = &a
fmt.Println(a,b)
*aPointer =  20
fmt.Println(a,b,aPointer,*aPointer)


}

