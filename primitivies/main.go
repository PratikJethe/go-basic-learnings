package main

import "fmt"

//integers
var (

	//booleans
	initializedBoolean   bool = false
	unInitializedBoolean bool

	//signed integers
	signed8BitInt  int8
	signed16BitInt int16
	signed32BitInt int32
	signed64BitInt int64

	//unsigned integers
	unsigned8BitInt  uint8
	unsigned16BitInt uint16
	unsigned32BitInt uint32

	//float
	float32Bit float32 = 3.14
	float64Bit float64 = 2.1e14

	//complex numbers
	complex64Bit complex64 = 1 + 2i
	complex32Bit complex128
)

func main() {

	a := 10 //1010
	b := 3  //0011

	fmt.Println(a & b)  //0010 AND
	fmt.Println(a | b)  //1011 OR
	fmt.Println(a ^ b)  //1001 EXOR
	fmt.Println(a &^ b) //0100 NOR

	fmt.Println(a << 3) // bit shift left 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // bit shift right 2^3 / 2^3 = 2^0
	fmt.Println(real(complex64Bit), imag(complex64Bit))

	s := "this is a string"
	s2 := " second string"
	byteValue := []byte(s)
	fmt.Println(s[1])         // prints byte valye
	fmt.Println(string(s[1])) // prints string value
	fmt.Println(s + s2)       // prints string value
	fmt.Println(byteValue)    // prints string value

	// runes are utf32
	// any utf 8 is valid utf32 character
	var r rune = 'r'
	fmt.Println(r)

}
