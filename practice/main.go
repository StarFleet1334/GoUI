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

	//var (
	//	n   int
	//	err error
	//)
	//
	//if n, err := strconv.Atoi(os.Args[1]); err == nil {
	//	fmt.Println("There was no error, n is:", n)
	//} else if err != nil {
	//	fmt.Println("There was an error:", err)
	//}1
	//
	//fmt.Printf("Because of Shadowing values are n: %d, err: %v", n, err)
	var (
		number int
		err    error
	)

	if n := os.Args; len(n) < 1 {
		fmt.Println("There was no argument")
	} else if err != nil {
		fmt.Println("There was an error")
	} else {
		number, err = strconv.Atoi(n[1])

	}

	fmt.Println(len(os.Args))
	fmt.Printf("%s", "X")
	for i := 0; i <= number; i++ {
		fmt.Printf("%10d", i)
	}
	fmt.Println()
	for i := 0; i <= number; i++ {
		fmt.Printf("%d", i)
		for j := 0; j <= number; j++ {
			fmt.Printf("%10d", i*j)

		}
		fmt.Println()

	}
}
