package main

import "fmt"

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
	fmt

}

func panicker() {
	fmt.Println("Create Panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error : ", err)
		}
	}()
	panic("Chia cho mot so la 0")
}
