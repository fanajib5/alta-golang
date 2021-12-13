package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// deklarasi struct
	// cara pertama
	var Person1 Student
	Person1.FirstName = "Faiq"
	Person1.LastName = "Najib"
	Person1.Age = 24
	fmt.Println(Person1)

	// cara kedua
	// urutan firstname dan lastname nya bisa ditukar
	var Person2 = Student{
		LastName:  "Najib2",
		FirstName: "Faiq",
		Age:       24,
	}
	fmt.Println(Person2)

	// cara ketiga
	Person3 := Student{"Faiq", "Najib3", 24}
	fmt.Println(Person3)
}
