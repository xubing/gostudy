package main

import (
	"encoding/json"
	"fmt"
	//	"strings"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func (mv Movie) I2MV(in *interface{}) (mvout Movie) {
	var ret = Movie{}
	//	ret.Actors = (*in).Actors
	//	ret.Year = (*in).released
	//	ret.Color = (*in).color
	//	ret.Actors =( *in).Actors
	return ret
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//解析

	book := []byte(`{"Title":"Go lang","released":"2014","color":true,"Actors":{"Age":20,"Name":"DD"}}`)
	var bookun Movie
	err = json.Unmarshal(book, &bookun)
	//	var mv = I2MV(book)

	fmt.Println(string(book))
}
