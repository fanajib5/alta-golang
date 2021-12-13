package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type planeFigure struct {
	Request  RequestBody  `json:"Request"`
	Response ResponseBody `json:"Response"`
}

type RequestBody struct {
	Base   string `json:"base"`
	Height string `json:"height"`
	Type   string `json:"type"`
}

type ResponseBody struct {
	Area string `json:"area"`
}

func (p *planeFigure) calculateArea(base int, height int, typePlaneFigure string) string {
	planeArea := ""
	planeAreaInt := 0

	switch typePlaneFigure {
	case "square":
		if base != height {
			planeArea = "Square must have same value either the base or the height"
		} else {
			planeAreaInt = base * height
		}
	case "rectangle":
		planeAreaInt = base * height
	case "triangle":
		baseFloat := float64(base)
		heightFloat := float64(height)
		var planeAreaFloat float64 = baseFloat * heightFloat / 2
		planeArea = strconv.FormatFloat(planeAreaFloat, 'f', 2, 64)
	case "parallelogram":
		planeAreaInt = base * height
	default:
		planeArea = "Please specify the type of plane figure"
	}

	if planeAreaInt != 0 {
		planeArea = strconv.Itoa(planeAreaInt)
	}

	return planeArea
}

func strToInt(str string) int {
	resInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return resInt
}

func apiArea(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	base := r.PostFormValue("base")
	height := r.PostFormValue("height")
	typeArea := r.PostFormValue("type")

	baseInt := strToInt(base)
	heightInt := strToInt(height)

	var p planeFigure
	areaVal := p.calculateArea(baseInt, heightInt, typeArea)
	p = planeFigure{
		Request: RequestBody{
			Base:   base,
			Height: height,
			Type:   typeArea,
		},
		Response: ResponseBody{
			Area: areaVal,
		},
	}

	responseBytes := new(bytes.Buffer)
	json.NewEncoder(responseBytes).Encode(p)

	w.Write(responseBytes.Bytes())
	fmt.Println("Endpoint Hit: calculate-area")
}

func main() {
	http.HandleFunc("/calculate-area", apiArea)
	fmt.Println("REST API serve at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
