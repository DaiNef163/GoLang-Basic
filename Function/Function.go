package main

import (
	"fmt"
)

func calc(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff

}

func main() {
	sum, diff := calc(15, 10)
	fmt.Println("Sum : ", sum)
	fmt.Println("Diff : ", diff)

}
