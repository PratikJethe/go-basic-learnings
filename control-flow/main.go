package main

import "fmt"

func main() {
	conMap := map[string]bool{
		"condition": true,
	}
	if conMap["condition"] {
		fmt.Println("executed")
	}

	if value, ok := conMap["condition"]; ok {
		fmt.Println(value)
	}

	number := 3
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println(number)
	}

	switch {
	case number > 5:
		fmt.Println("number greater than 5")
	case number < 5:
		fmt.Println("number less than 5")
	}
	switch {
	case number < 5:
		fmt.Println("number less than 5")
		fallthrough // forced to exceute next case
	case number > 5:
		fmt.Println("number greater than 5")
	}


	var any interface {} =1.4
	switch any.(type) {
	case int:
		fmt.Println("int type")
	case float32 , float64:
		fmt.Println("float type")
		break
		fmt.Println("dont print this")
	}
}
