package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	//age := os.Args[1]
	//n, err := strconv.Atoi(age)
	//
	//if err != nil {
	//	fmt.Printf("There was a error %v", err)
	//	return
	//} else {
	//	fmt.Printf("There was no error, answer was %v", n)
	//}
	//

	// Simple Statement usage

	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Println("There was no error, n is:", n)
	}
}
