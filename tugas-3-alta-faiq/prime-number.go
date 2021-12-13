package main

import "fmt"

func main() {
	var myNum int
	fmt.Println("Program untuk mencari bilangan kelipatan 5 dan bilangan primanya")
	fmt.Printf("Masukkan bilangan kelipatan '5' : ")
	fmt.Scanf("%d", &myNum)

	if myNum%5 == 0 {
		fmt.Printf("Bilangan %d adalah kelipatan 5\n", myNum)
		var primeNum []int
		var sumPrime int

		for i := 2; i < myNum; i++ {
			var isPrime bool = true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime == true {
				primeNum = append(primeNum, i)
				sumPrime += i
			}
		}
		fmt.Println("Bilangan primanya:", primeNum)
		fmt.Println("Jumlah dari bilangan primanya =", sumPrime)
	} else {
		fmt.Printf("Bilangan %d BUKAN kelipatan 5", myNum)
	}
}
