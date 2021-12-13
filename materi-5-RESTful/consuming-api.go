package main

import (
	"fmt"
	"net/http"
	"os"
)

// A response struct to map the Entire Response
type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
}

func main() {
	// Request data to API
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(response)

	// // read data from response
	// responseData, _ := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()

	// // Parsing data to Data Structure
	// var LukeSkywalker StarWarsPeople
	// json.Unmarshal(responseData, &LukeSkywalker)

	// fmt.Println("------ Print Field ------")
	// fmt.Println(LukeSkywalker.Name)
	// fmt.Println(LukeSkywalker.Height)
	// fmt.Println(LukeSkywalker.Mass)
	// fmt.Println(LukeSkywalker.HairColor)
	// fmt.Println(LukeSkywalker.Films[0])
}
