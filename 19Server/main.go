package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server setup")
	GetReq()
}

func GetReq(myurl string) {
	givenUrl := myurl
	res, err := http.Get(givenUrl)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
}
