package main

import "fmt"

func calculateFactorial(n int) (int, []int) {
	var factorialNumbers []int
	factorialResult := 1

	for i := 1; i <= n; i++ {
		factorialNumbers = append(factorialNumbers, i)
		factorialResult *= i
	}

	return factorialResult, factorialNumbers
}

func printFactorial(inputNumber int) {
	var factorialForm []int
	var factorialNumber int
	factorialNumber, factorialForm = calculateFactorial(inputNumber)

	fmt.Printf("%d! = ", inputNumber)
	for index := 0; index < len(factorialForm); index++ {
		if index == (len(factorialForm) - 1) {
			fmt.Printf("%d", factorialForm[index])
		} else {
			fmt.Printf("%d * ", factorialForm[index])
		}
	}
	fmt.Printf(" = %d", factorialNumber)
}

func main() {
	var inputNumber int
	fmt.Println("Program untuk menghitung nilai faktorial dari bilangan-n")
	fmt.Printf("Masukkan nilai bilangan-n: ")
	fmt.Scanf("%d", &inputNumber)

	fmt.Println("--------------------------------")
	printFactorial(inputNumber)
}
