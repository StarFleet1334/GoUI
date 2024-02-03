package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesA_slice := strings.Split(namesA, ",")
	slices.Sort(namesA_slice)
	slices.Sort(namesB)

	if len(namesA_slice) != len(namesB) {
		fmt.Println("Two slices are not equal")
		return
	} else {
		for index, _ := range namesA_slice {
			if namesA_slice[index] != namesB[index] {
				fmt.Println("Two slices are not equal")
				return
			}
		}
		fmt.Println("Two slices are equal")
	}
}
