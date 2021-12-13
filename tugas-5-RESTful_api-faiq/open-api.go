package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type ThroneChara struct {
	ID        int    `json:"id"`
	FristName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Family    string `json:"family"`
	Image     string `json:"image"`
	ImageUrl  string `json:"imageUrl"`
}

func getThroneChara(response *http.Response) []byte {
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBodyBytes
}

func main() {
	response, err := http.Get("https://thronesapi.com/api/v2/Characters/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	byteChara := getThroneChara(response)

	// Convert response body byte to struct
	var charaDetails ThroneChara
	json.Unmarshal(byteChara, &charaDetails)
	fmt.Printf("JSON Response in struct: %+v\n", charaDetails)
}
