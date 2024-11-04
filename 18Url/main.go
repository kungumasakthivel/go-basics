package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com:3000/dev?query=hello"

func main() {
	fmt.Println("URL in go")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("query param type is %T\n", qparams)

	fmt.Println(qparams["query"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/dev",
		RawQuery: "user=manish",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
