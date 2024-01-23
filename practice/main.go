package main

import (
	"fmt"
	"os"
)

func main() {

	// Getting Input From Terminal using go
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("First argument:", os.Args[0])
	fmt.Print("Second argument:", os.Args[1])

	fmt.Print("Total number of arguments:", len(os.Args))
}
