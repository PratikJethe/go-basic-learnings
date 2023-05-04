package main

import (
	"fmt"
	"reflect"
)

func main() {

	//MAPS
	numberMap := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(numberMap)
	// numberMap = make(map[string]int) //using make
	delete(numberMap, "one")
	fmt.Println(numberMap)
	//ok comma syntax
	element, ok := numberMap["one"]
	println(element, ok)
	println(len(numberMap))

	numberMapCopy := numberMap //copied by reference

	numberMapCopy["three"] = 3

	fmt.Println(numberMap)
	fmt.Println(numberMapCopy)

	//STRUCTS

	type Person struct {
		name       string
		PublicName string
		age        int
		hobbies    []string
	}

	newPerson := Person{
		name:       "pratik",
		PublicName: "public name",
		age:        23,
		hobbies:    []string{"reading", "music", "piano"},
	}
	shortLivedStruct := struct {
		name string
	}{name: "name"}
	fmt.Println(newPerson)
	fmt.Println(newPerson.name)
	fmt.Println(shortLivedStruct)


	type Animal struct {
		Name string `required max:"100"`
		Origin string
	}

	type Bird struct {
		Animal  // embedding structs like inheritence
		SpeedKPH float32
		
	}

	b:=Bird{}
	b.Name ="EMU"
	b.Origin="Australia"
    b.SpeedKPH = 52

	fmt.Println(b)

	// stuct tags
	t:=reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
