package main

import "fmt"

type Car struct {
	Name  string
	Model string
	Year  int
	Color string
}

func printCar(c Car) {
	fmt.Println("Mobil\t\t:", c.Name)
	fmt.Println("Model\t\t:", c.Model)
	fmt.Println("Tahun produksi\t:", c.Year)
	fmt.Println("Warna\t\t:", c.Color)
}

func (c Car) newCar() string {
	printCar(c)
	return "New Released!!"
}

func (c Car) legendaryCar(oldModel string, year int) string {
	c.Model = oldModel
	c.Year = year
	printCar(c)
	return "Mobil Legendaris~"
}

func main() {
	c := Car{
		Name:  "Subaru",
		Model: "WRX STI",
		Year:  2019,
		Color: "Blue",
	}

	fmt.Printf("Struct mobil\n\n")
	fmt.Println(c.newCar())
	fmt.Println("--------------------")
	oldModel := "Impreza WRX"
	oldYear := 1992
	legend := c.legendaryCar(oldModel, oldYear)
	fmt.Println(legend)
}
