package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	unmarshalDemo2()
}

func marshalDemo1() {
	movies := Movie{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	}
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func unmarshalDemo1() {
	var jsonStr = `{
	"attendance_batch_id":35,
	"code":"ATT00004",
	"attendance_cost":"2.75",
	"time":"2023-04-07至2023-04-08"
}`
	var jsonObject map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObject)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Printf("%+v", jsonObject)
}

type AnimalData struct {
	Bird int    `json:"bird"`
	Cat  string `json:"cat"`
}

func unmarshalDemo2() {
	text := "{\"bird\":10,\"cat\":\"Fuzzy\"}"
	var animal AnimalData
	json.Unmarshal([]byte(text), &animal)
	// Print result.
	fmt.Printf("BIRD = %v, CAT = %v", animal.Bird, animal.Cat)
}
