package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://api.restful-api.dev/objects?id=3&name=Rajath"

func main() {
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Println(qParams)

	fmt.Println(qParams["id"][0])
	fmt.Println(qParams["name"][0])

}
