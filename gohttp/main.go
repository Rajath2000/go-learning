package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://api.restful-api.dev/objects"

func main() {
	res, err := http.Get(Url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.Body)

	resultInBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(resultInBytes))
}
