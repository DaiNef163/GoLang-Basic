package main

import (
	"fmt"
)

func nhapmang(n int, arr *[]int) {
	fmt.Println("Nhap so luong mang")
	fmt.Scan(&n)
	*arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("a[", i, "] = ")
		fmt.Scan(&(*arr)[i])
	}
}

func inmang(arr []int) {
	fmt.Println("--------------------")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i] , " ")
	}
	fmt.Println()
}

func timX(arr []int, x int) bool {
	fmt.Println("Nhap x")
	fmt.Scan(&x)

	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func hoanvi(arr []int, i, j int) {
	var temp int = arr[i]
	arr[i] = arr[j]
	arr[j] = temp

}
func sapxep(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				hoanvi(arr, i, j)
			}

		}
	}

}

func main() {
	var n int
	var arr []int
	var x int
	nhapmang(n, &arr)
	inmang(arr)
	fmt.Println(timX(arr, x))
	sapxep(arr)
	inmang(arr)
	
}
