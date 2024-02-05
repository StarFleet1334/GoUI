package main

import (
	"fmt"
	"slices"
)

type collection []string

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	slices.Sort(items[(len(items)/2)-1 : (len(items)/2)+2])
	fmt.Println()
	fmt.Println("Sorted  :", items)

}
