package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tage     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in golang")
	EncodeJson()
}

func EncodeJson() {
	c := []course{
		{"Reactjs", 299, "YT", "1387439kd", []string{"web", "dev", "backend"}},
		{"NodeJS", 349, "YT", "1389kd", []string{"web", "dev", "backend"}},
		{"NodeJS", 349, "YT", "1389kd", nil},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}
