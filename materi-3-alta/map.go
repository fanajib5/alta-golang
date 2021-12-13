package main

import (
	"fmt"
	"strings"
)

func main() {
	// map
	// var harga = map[string]int{"siomay": 1000, "batagor": 500}
	// harga := map[string]int{"siomay": 1000, "batagor": 500}
	var harga = make(map[string]int)
	harga["bakwan"] = 7000
	harga["siomay"] = 7000
	delete(harga, "bakwan")
	fmt.Println(harga["bakwan"])

	fmt.Println(strings.Index("alterra", "ra"))
}
