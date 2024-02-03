package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		pizza_toppings   []string
		departure_times  []time.Time
		graduation_years []int
		states_of_lights []bool
	)

	pizza_toppings = append(pizza_toppings, "pepperoni", "onions", "extra cheese")

	now := time.Now()
	departure_times = append(departure_times,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	graduation_years = append(graduation_years, 1998, 2005, 2018)

	states_of_lights = append(states_of_lights, true, false, true)

	fmt.Printf("pizza       : %s\n", pizza_toppings)
	fmt.Printf("\ngraduations : %d\n", graduation_years)
	fmt.Printf("\ndepartures  : %s\n", departure_times)
	fmt.Printf("\nlights      : %t\n", states_of_lights)
}
