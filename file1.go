package main

import (
	"fmt"
)

type Student struct {
	name     string
	age      int
	class    string
	subjects []string
}

type Car struct {
	name    string
	price   int
	color   string
	country []string
}

func main() {

	Car1 := Car{
		name:  "Maybach 680",
		price: 123456789,
		color: "black",
		country: []string{
			"Viet Nam ,",
			"Lao",
		},
	}

	Student1 := Student{
		name:  "Nam",
		age:   20,
		class: "ITA",
		subjects: []string{
			"Math ,",
			"Science ,",
			"English",
		},
	}

	fmt.Println(Car1)
	fmt.Println(Student1)

	fmt.Println()

	// var nameMap map[keytype] valueType
	//keytype : kieu du lieu cua key trong map
	//valuetype : la kieu du lieu cua gia tri tuong ung moi key

	// var map1 map[string]string
	// myMap = make(map[keyType]valueType)
	nameStudents := make(map[string]string)
	ageStudents := make(map[string]int)
	// them phan tu vao map

	// nameMap[key] = value

	nameStudents["Student1"] = "Nguyen Van A"
	nameStudents["Student2"] = "Nguyen Van B"
	nameStudents["Student3"] = "Nguyen Van C"
	nameStudents["Student4"] = "Nguyen Van D"
	nameStudents["Student5"] = "Nguyen Van E"
	nameStudents["Student6"] = "Nguyen Van F"
	nameStudents["Student7"] = "Nguyen Van G"
	nameStudents["Student8"] = "Nguyen Van H"
	nameStudents["Student9"] = "Nguyen Van I"
	nameStudents["Student10"] = "Nguyen Van J"

	ageStudents["Student1"] = 18
	ageStudents["Student2"] = 20
	ageStudents["Student3"] = 12
	ageStudents["Student4"] = 14
	ageStudents["Student5"] = 16

	chooseNameStudent := nameStudents["Student2"]
	chooseAgeStudent := ageStudents["Student2"]

	fmt.Println("Student 2 ->  Name : ", chooseNameStudent, " , Age :", chooseAgeStudent)

	fmt.Println()

	keymap := make(map[string]int)
	keymap["key 1"] = 1
	choosekey := keymap["key 1"]
	fmt.Println(choosekey)

	fmt.Println()

}
