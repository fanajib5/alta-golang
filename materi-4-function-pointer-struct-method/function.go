package main

import (
	"fmt"
	"math"
)

// contoh function tanpa parameter
func sayHello() {
	fmt.Println("Hello")
}

// contoh function dengan parameter
func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi!")
	} else if hour < 18 {
		fmt.Println("Selamat Sore!")
	} else {
		fmt.Println("Selamat Malam!")
	}
}

// contoh function dengan single return value
func calculateSquare(side int) int {
	return side * side
}

// contoh function dengan multiple return value
func calculateCircle(diameter float64) (float64, float64) {
	var luas = math.Pi * math.Pow(diameter/2, 2)
	var keliling = math.Pi * diameter

	return keliling, luas
}

func main() {
	// declare variabel dulu,
	// biar ngga jadi magic number
	hour := 13

	sayHello()
	greeting(hour)

	// persegi dan lingkaran
	side := 5
	wide := calculateSquare(side)
	fmt.Printf("Luas persegi= %d\n\n", wide)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)

	fmt.Printf("Luas lingkaran= %.2f\n", luas)
	fmt.Printf("Keliling lingkaran= %.2f\n", keliling)
}
