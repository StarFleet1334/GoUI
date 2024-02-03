package main

import "fmt"

func main() {

	// Declaration of Arrays

	names := [3]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.14145}
	alives := [...]bool{true, false, true, false}
	zero := [0]byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	names_slices := []string{"Einstein", "Tesla", "Shepard"}
	distances_slices := []int{50, 40, 75, 30, 125}
	data_slices := []byte{'H', 'E', 'L', 'L', 'O'}
	ratios_slices := []float64{3.14145}
	alives_slices := []bool{true, false, true, false}
	zero_slices := []byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names_slices)
	fmt.Printf("distances: %[1]T %[1]d\n", distances_slices)
	fmt.Printf("data     : %[1]T %[1]d\n", data_slices)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios_slices)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives_slices)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero_slices)

}
