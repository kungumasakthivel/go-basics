package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Server setup")
	// GetReq("https://zuai-tkx1.onrender.com/posts")
	// PostReq("https://zuai-tkx1.onrender.com/posts")
	PostFormReq()
}

func GetReq(myurl string) {
	givenUrl := myurl
	res, err := http.Get(givenUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	fmt.Println(string(content))
}

func PostReq(myurl string) {

	reqBody := strings.NewReader(`
		{
			"title":"go",
			"body": "golang"
		}
	`)

	res, _ := http.Post(myurl, "application/json", reqBody)
	res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PostFormReq() {
	const murl = "https://zuai-tkx1.onrender.com/posts"

	data := url.Values{}
	data.Add("title", "go title")
	data.Add("body", "go body")

	res, err := http.PostForm(murl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
