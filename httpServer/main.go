package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const Url = "https://jsonplaceholder.typicode.com/comments"

type Comments struct {
	PostId string `json:"posId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	res, err := http.Get(Url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.Body)

	DecodeJsonToGo(res.Body)

	// resultInBytes, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// var resultString strings.Builder

	// count, _ := resultString.Write(resultInBytes)

	// fmt.Println(count)

	// fmt.Println(resultString.String())

	//how to handle JSON
	// EncodeJsonData()
}

func DecodeJsonToGo(response io.Reader) {
	var comments []Comments

	// resultInBytes, err := io.ReadAll(response)

	// if err != nil {
	// 	panic(err)
	// }

	// checkValid := json.Valid(resultInBytes)

	if checkValid := true; checkValid {

		err := json.NewDecoder(response).Decode(&comments)

		if err != nil {
			panic(err)
		}

		for _, comment := range comments {
			fmt.Println(comment.Id)
		}
	}
}

func EncodeJsonData() {

	comments := []Comments{
		{"1", 1, "Rajath M R", "rajatmr2000@gmail.com", "Hello"},
		{"1", 1, "Rajath M R", "rajatmr2000@gmail.com", "Hello"},
		{"1", 1, "Rajath M R", "rajatmr2000@gmail.com", "Hello"},
		{"1", 1, "Rajath M R", "rajatmr2000@gmail.com", "Hello"},
	}

	data, err := json.MarshalIndent(comments, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
