package main

import "fmt"

func main() {

	// Declaring nil Slices
	var names []string

	var distances []int

	var data []uint8

	var ratios []float64

	var alives []bool

	fmt.Printf("names    : %#v %d %t \n", names, len(names), names == nil)

	fmt.Printf("distances    :  %#v %d %t \n", distances, len(distances), data == nil)

	fmt.Printf("data    :  %#v %d %t \n", data, len(data), data == nil)

	fmt.Printf("ratios    :  %#v %d %t \n", ratios, len(ratios), ratios == nil)

	fmt.Printf("alives    :  %#v %d %t", alives, len(alives), alives == nil)

}
