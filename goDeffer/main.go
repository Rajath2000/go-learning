package main

import "fmt"

func main() {

	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	defer mydefer()

}

func mydefer() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
