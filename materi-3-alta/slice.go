package main

import "fmt"

func main() {
	// Slice A={1,2,3,4,5}  maka len = 5  dan  cap = 5
	// ditambahkan elemen, maka Slice A={1,2,3,4,5,6}  maka  len = 6  dan  cap = 10
	// karena capacity akan diatur golang ( 2 * cap_awal )

	// var evenNumbers []int //  <--- inilah yg disebut slice, tetapi tanpa value
	// var oddNumbers []int{1,3,5,7}  // <--- slice dengan nilai

	// deklarasi slice dengan keyword make
	// var primes = make([]int, 5, 10)

	// APPEND & COPY
	var colors = []string{"red", "green", "blue"}
	colors = append(colors, "purple") // <--- menambah value pada slice

	copiedColors := make([]string, 10)
	copy(copiedColors, colors)
	fmt.Println(copiedColors)
}
