package main

import (
	"fmt"
	"sync"
)

var sampleArray []int

var wg sync.WaitGroup

var mut sync.Mutex

func main() {

	wg.Add(3)

	go fucntion1()

	go fucntion2()

	go fucntion3()

	wg.Wait()
	fmt.Println(sampleArray)
}

func fucntion1() {
	fmt.Println("This is in function 1")

	mut.Lock()
	sampleArray = append(sampleArray, 1)
	mut.Unlock()
	wg.Done()

}

func fucntion2() {
	fmt.Println("This is in function 2")

	mut.Lock()
	sampleArray = append(sampleArray, 2)
	mut.Unlock()
	wg.Done()
}

func fucntion3() {

	fmt.Println("This is in function 3")

	mut.Lock()
	sampleArray = append(sampleArray, 3)
	mut.Unlock()
	wg.Done()

}
