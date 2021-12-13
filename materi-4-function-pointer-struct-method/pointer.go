package main

import "fmt"

func main() {
	var name string = "John"
	var nameAddress *string = &name
	fmt.Println("name (value) :", name)
	fmt.Println("name (address) :", &name)
	fmt.Println("nameAddress (value) :", *nameAddress)
	fmt.Println("nameAddress (address of name) :", nameAddress)
	fmt.Println("nameAddress (address) :", &nameAddress)

	name = "Doe"
	fmt.Println("name (value) :", name)
	fmt.Println("name (address) :", &name)
	fmt.Println("nameAddress (value) :", *nameAddress)
	fmt.Println("nameAddress (address of name) :", nameAddress)
	fmt.Println("nameAddress (address) :", &nameAddress)
}
