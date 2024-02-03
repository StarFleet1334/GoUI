package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	arg := os.Args

	if len(arg) <= 1 {
		fmt.Println("provide a few numbers")
	} else {
		var lst []int
		arr := arg[1:]

		for i := 0; i < len(arr); i++ {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				continue
			}
			lst = append(lst, num)
		}
		slices.Sort(lst)
		fmt.Println(lst)
	}
}
