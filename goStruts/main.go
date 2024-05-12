package main

import "fmt"

func main() {
	rajath := User{"Rajath", "rajathmr2000@gmail.com", true, 24}
	fmt.Println(rajath)

	if rajath.Age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
