package main

import "fmt"

func main() {
	fmt.Println("Algoritma Cuci Piring")
	// banyak piring kotor
	piringKotor := 10

	for piring := 1; piring <= piringKotor; piring++ {
		fmt.Printf("Piring ke-%d\n", piring)
		tetesSabun := 1
		for true {
			tetesSabun++
		}
		fmt.Printf(" membutuhkan %d tetes sabun cuci piring.", tetesSabun)
	}
}
