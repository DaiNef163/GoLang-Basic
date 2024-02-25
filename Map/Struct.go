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
type MonHoc struct {
	chuyencan        string
	giuaky           string
	cuoiky           string
	buoihoctrongtuan []string
}

func main() {

	Mon1 := MonHoc{
		chuyencan: "100",
		giuaky:    "100",
		cuoiky:    "100",
		buoihoctrongtuan: []string{
			"Thu hai",
			"Thu ba ",
		},
	}

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
	Student2 := Student{
		name:  "Nguyen Hoai Nam",
		age:   20,
		class: "TPM10",
		subjects: []string{
			"C++ ",
			"Java , ",
			"GoLang",
		},
	}

	fmt.Println(Car1)
	fmt.Println(Student1)
	fmt.Println(Student2)
	fmt.Println(Mon1)
	fmt.Println()
}
