package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

// function dalam bentuk method
func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {
	e := Employee{
		FirstName: "Faiq",
		LastName:  "Najib",
	}

	fmt.Println(e.fullName())
}
