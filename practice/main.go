package main

import "fmt"

func main() {

	// Declaring nil Slices

	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alives    []bool
	)
	names = []string{"Saba", "Luka", "Gio"}
	distances = []int{1, 2, 3, 4, 5}
	data = []uint8{'I', 'N', 'A', 'N', 'C'}
	ratios = []float64{1.1, 1.2}
	alives = []bool{true, false, false, false}

	fmt.Printf("names    : %T %d %t \n", names, len(names), names == nil)

	fmt.Printf("distances    :  %T %d %t \n", distances, len(distances), data == nil)

	fmt.Printf("data    :  %T %d %t \n", data, len(data), data == nil)

	fmt.Printf("ratios    :  %T %d %t \n", ratios, len(ratios), ratios == nil)

	fmt.Printf("alives    :  %T %d %t", alives, len(alives), alives == nil)

}
