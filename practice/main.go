package main

import (
	"fmt"
	"github.com/inancgumus/prettyslice"
)

func main() {

	//header := []string{"Location", "Size", "Beds", "Baths", "Price"}
	//type placeholder [5]string
	//
	//var data []placeholder
	//
	//data = append(data, placeholder{
	//	"New York", "100", "2", "1", "100000",
	//}, placeholder{
	//	"New York", "150", "3", "2", "200000",
	//}, placeholder{
	//	"Paris", "200", "4", "3", "400000",
	//}, placeholder{
	//	"Istanbul", "500", "10", "5", "100000",
	//})
	//
	//for _, item := range header {
	//	fmt.Printf("%-15s", item)
	//}
	//fmt.Println("\n===========================================================================")
	//var result [5]int
	//
	//for _, item := range data {
	//	for i, insideItem := range item {
	//		if i > 0 {
	//			num, _ := strconv.Atoi(insideItem)
	//			result[i] = result[i] + (num)
	//		}
	//		fmt.Printf("%-15s", insideItem)
	//	}
	//	fmt.Println("\n")
	//}
	//fmt.Println("\n===========================================================================")
	//for i, item := range result {
	//	if i == 0 {
	//		fmt.Printf("%-15s", "")
	//
	//	} else {
	//		fmt.Printf("%-15.2f", float64(item)/float64(len(data[0])))
	//	}
	//}

	ages := []int{35, 15, 25}
	red, green := ages[0:1], ages[1:3]

	prettyslice.Show("ages", ages)
	prettyslice.Show("red", red)
	prettyslice.Show("green", green)

	fmt.Println(red[0])
	// fmt.Println(red[1]) // error
	// fmt.Println(red[2]) // error

	{
		var ages []int
		prettyslice.Show("nil slice", ages)

		// or just:
		prettyslice.Show("nil slice", []int(nil))
	}
}
