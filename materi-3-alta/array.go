package main

import (
	"fmt"
)

func main() {
	// var primes int
	// var countries [5]string

	// fmt.Println(reflect.ValueOf(primes).Kind())
	// fmt.Println(reflect.ValueOf(countries).Kind())

	// arr := [5]int{1, 3, 5, 7, 9}
	// var arrPartial [5]int = [5]int{0, 2, 4}

	// fmt.Println(arr)
	// fmt.Println(arrPartial)

	// ---------------------------
	// LOOPING IN ARRAY
	primes := [5]int{2, 1, 5}
	// Cara loop pertama
	for index := 0; index < len(primes); index++ {
		fmt.Println(primes[index])
	}

	// Cara loop kedua
	for index, element := range primes {
		fmt.Println(index, " => ", element)
	}

	// Cara loop ketiga
	index := 0
	for range primes {
		fmt.Println(primes[index])
		index++
	}
}
