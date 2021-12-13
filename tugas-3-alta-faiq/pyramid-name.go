package main

import "fmt"

func main() {
	myName := "Faiq"
	for i := 0; i < len(myName); i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(string(myName[j]) + " ")
		}
		fmt.Println("")
	}
	for i := len(myName); i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf(string(myName[j]) + " ")
		}
		fmt.Println("")
	}
}
