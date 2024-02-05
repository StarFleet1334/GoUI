package main

import (
	"fmt"
	"github.com/inancgumus/prettyslice"
)

type collection []string

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	games := []string{}
	//
	// 2. print the length and capacity of the games slice
	//fmt.Printf("Length of games slice: %d\n", len(games))
	//fmt.Printf("Capacity of games slice: %d\n", cap(games))

	//
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	// 4. print the length and capacity of the games slice
	//
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	games = append(games, "pacman", "mario", "tetris", "doom")

	//
	// 6. print the length and capacity of the games slice
	//fmt.Printf("Length of games slice: %d\n", len(games))
	//fmt.Printf("Capacity of games slice: %d\n", cap(games))
	//
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	for i := 0; i <= len(games); i++ {
		s := games[:i]
		prettyslice.Show("After", s)
	}

	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	// for ... {
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", ...)
	// }

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]
	//
	// 2. print the games and the new slice's len and cap
	prettyslice.Show("Zero", zero)

	//
	// 3. append a new element to the new slice
	//
	// 4. print the new slice's lens and caps
	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		prettyslice.Show("Zero (updated)", zero)
	}

	prettyslice.Show("Games", games)
	//
	// 5. repeat the last two steps 5 times (use a loop)
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	// []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	fmt.Println()

	// zero := ...
	// fmt.Printf("games's     len: %d cap: %d\n", ...)
	// fmt.Printf("zero's      len: %d cap: %d\n", ...)

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	// zero = ...
	// for ... {
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", ...)
	// }

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()

	fmt.Println()

	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity

	// --- #7 ---
	// uncomment and see the error.
	_ = games[:cap(games)+1]
	//or:
	_ = games[:5]

}
