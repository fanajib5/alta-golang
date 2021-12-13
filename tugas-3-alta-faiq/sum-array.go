package main

import "fmt"

func main() {
	numbers := [5]int{1, 3, 5, 7, 9}
	fmt.Printf("Array of numbers = %d\n", numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Jumlah = %d", sum)
}
