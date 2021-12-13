package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array of numbers = %d\n", numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Jumlah = %d\n\n", sum)

	// myName := "Faiq"
	// for i := 0; i < len(myName); i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Printf(string(myName[j]) + " ")
	// 	}
	// 	fmt.Println("")
	// }
	// for i := len(myName); i > 0; i-- {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Printf(string(myName[j]) + " ")
	// 	}
	// 	fmt.Println("")
	// }

	var myNum int
	fmt.Printf("Input the n number: ")
	fmt.Scanf("%d", &myNum)
	var primeNum []int

	for i := 2; i < myNum; i++ {
		var isPrime bool
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = true
				break
			}
		}
		if isPrime == false {
			primeNum = append(primeNum, i)
		}
	}
	fmt.Println(primeNum)
}
