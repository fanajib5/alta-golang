package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piringKotor := 10
	fmt.Println("Ini adalah implementasi algoritma cuci piring :)")
	fmt.Printf("Kondisi saat ini diasumsikan memiliki %d piring kotor dan,\n", piringKotor)
	fmt.Printf("tiap piring memiliki tingkat noda yang berbeda-beda secara random.\n\n")
	for piring := 1; piring <= piringKotor; piring++ {
		tetesSabun := 1
		tingkatNoda := rand.Intn(6)
		fmt.Println("------------------------------------------------------------------------------------------")
		fmt.Printf("Piring ke-%d memiliki tingkat noda '%d'.\n", piring, tingkatNoda)
		for tetes := 1; tetes < 10; tetes++ {
			fmt.Printf("\t+ Piring ke-%d menggunakan %d tetes sabun.\n", piring, tetesSabun)
			fmt.Printf("Apakah piring ke-%d sudah bersih?\n", piring)
			if tingkatNoda == tetes || tingkatNoda == 0 {
				fmt.Printf("---> Yes! Piring ke-%d sudah bersih. Terima kasih.\n", piring)
				break
			} else {
				fmt.Printf("\t+ No! Piring ke-%d belum bersih, tambah 1 tetes sabun lagi untuk yang ke-%d kalinya.\n", piring, tetesSabun)
				tetesSabun++
			}
		}
		fmt.Printf("Piring ke-%d dengan tingkat noda '%d',  butuh %d tetes sabun\n", piring, tingkatNoda, tetesSabun)
		fmt.Print("------------------------------------------------------------------------------------------\n\n")
	}
}
