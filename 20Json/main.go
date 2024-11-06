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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "NodeJS",
			"price": 349,
			"website": "YT",
			"tags": [
					"web",
					"dev",
					"backend"
			]
        }
	`)

	// check if its json format is correct
	var data course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("data is valid")
		json.Unmarshal(jsonDataFromWeb, &data)
		fmt.Printf("%#v\n", data)
	} else {
		fmt.Println("Data is not valid")
	}

	// some cases where you want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", key, value)
		fmt.Printf("Value type %T\n", value)
	}
}
