package main

import (
	"fmt"	
)

type MyNumber struct {
	number int
}

func main() {
	var MyNumber1 *MyNumber
	var MyNumber2 *MyNumber
	fmt.Println(MyNumber2)
	MyNumber1 = &MyNumber{number: 10}
	MyNumber2 = new(MyNumber)

	fmt.Println(MyNumber1)
	fmt.Println(MyNumber2)

	(*MyNumber2).number = 99 //MyNumber2.number = 99
	fmt.Println(MyNumber2.number)

	var a = map[string]string{"Seed1": "GAM", "Seed2": "T1"}
	b := a
	fmt.Println(a, b)
	a["Seed2"] = "GenG"
	fmt.Println(a, b)

}
