package main

import "fmt"

func main() {
	// IF, ELSE IF, ELS
	hour := 20
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}

	// SWITCH
	var today int = 3
	switch today {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Invalid date")
	}
}
