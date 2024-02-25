package main

import "fmt"

func main() {
	var a int = 12
	var b *int = &a
	fmt.Println(a, *b)
	a = 24
	fmt.Println(a, *b)
	fmt.Println("----------")

	m := []int{1, 2, 3}
	n := m
	fmt.Println(m, n)

	m[1] = 99
	m[2] = 100
	fmt.Println(m, n)

}
