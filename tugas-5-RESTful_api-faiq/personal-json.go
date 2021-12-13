package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type PersonalData struct {
	Name         string   `json:"nama"`
	PlaceOfBirth string   `json:"tempat_lahir"`
	DateOfBirth  string   `json:"tanggal_lahir"`
	Height       int      `json:"height"`
	Mass         int      `json:"mass"`
	Phone        string   `json:"no_hp"`
	Email        string   `json:"email"`
	Address      string   `json:"alamat"`
	City         string   `json:"kota"`
	State        string   `json:"provinsi"`
	Country      string   `json:"negara"`
	PostalCode   int      `json:"kode_pos"`
	Gender       string   `json:"jenis_kelamin"`
	FavoriteFilm []string `json:"favorite_film"`
}

func (d PersonalData) prettyJson() string {
	jsonPrettyIndent, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonPrettyIndent)
}

func main() {
	d := PersonalData{
		Name:         "Faiq Najib Al-Aziz",
		PlaceOfBirth: "Jombang",
		DateOfBirth:  "07 Mei 1997",
		Height:       177,
		Mass:         68,
		Phone:        "082336071376",
		Email:        "faiq.najib@gmail.com",
		Address:      "Pagesangan Agung I",
		City:         "Surabaya",
		State:        "Jawa Timur",
		Country:      "Indonesia",
		PostalCode:   60233,
		Gender:       "Laki-Laki",
		FavoriteFilm: []string{"Iron Man", "3 Idiots"},
	}
	fmt.Println("Personal JSON Data\n", d.prettyJson())
}
