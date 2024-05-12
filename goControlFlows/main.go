package main

import "fmt"

func main() {

	// if else

	var resultString string

	percentage := 29

	if percentage < 30 {
		resultString = "Fail"
	} else if percentage >= 30 && percentage <= 50 {
		resultString = "pass"
	} else if percentage > 50 && percentage <= 85 {
		resultString = "firstClass"
	} else {
		resultString = "Distinction"
	}

	//switch case

	switch resultString {
	case "Fail":
		fmt.Println("You are Failed!")

	case "Pass":
		fmt.Println("Good job you are passed!")

	case "firstClass":
		fmt.Println("Well done you are first class!")

	default:
		fmt.Println("Exelent you are disntinction")
	}

	//loops

	days := []string{"Monday", "tuesday", "wednesday"}
	fmt.Println(days)

	for i, day := range days {
		fmt.Println(i, day)
	}

	//functions
	value := addNumber(3, 5)

	fmt.Println(value)

	//methods
	rajath := User{"Rajath", "rajathmr2000@gmail.com", true, 24}
	rajath.GetStaus()
	rajath.updateMail("rahul@gamil.com")
	rajath.getEmail()

}

func addNumber(number1 int, number2 int) int {
	return number1 + number2
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStaus() {
	fmt.Println(u.Status)
}

func (u User) updateMail(newEmail string) {
	u.Email = newEmail
}

func (u User) getEmail() {
	fmt.Println(u.Email)
}
