package main

import (
	"fmt"
)

func main() {
	// for i := 1; i < 10; i++ {
	// 	if i%2 == 0 {

	// 		fmt.Print(i, " ")
	// 	}
	// }

	// i := 1
	// for ; i<=10 ; i++ {
	// 	fmt.Print(i, " ")
	// }

	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= 9; j++ {
	// 		fmt.Println(i, " * ", j, " = ", i*j)

	// 	}
	// 	fmt.Println("=======================")
	// }

	// arr := [3]int{2, 3, 4}
	// for i := 0; i < len(arr); i++ {

	// 	fmt.Print(arr[i], " ")
	// }

	// arr := []int{2, 3, 4}
	// for index, value := range arr {

	// 	fmt.Print("a[", index, "]")
	// 	fmt.Print(value)
	// 	fmt.Println()
	// }

	s := "Hello World"
	for index, value := range s {
		fmt.Print("a[", index, "] ")
		fmt.Print(string(value))
		fmt.Println()
		fmt.Println()
	}
}
