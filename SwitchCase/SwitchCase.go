package main

import "fmt"

func main() {

	var a int
	fmt.Println("Input a")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Printf("key 1")
	case 2:
		fmt.Println("key 2")

	default:
		fmt.Printf("Error")
	}
}
