package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 4; i < 5; i, j = i+1, j-1 {

		// if i ==2  && j ==2 {
		// 	continue
		// }
		if i == 2 && j == 2 {
			break
		}
		fmt.Println(i, j)
	}

	// loop with onlu condition like while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++

	}

	//infinite loop
	// for {

	// }

	// range loop
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for index, v := range s {
		fmt.Println(index, v)
	}
}
