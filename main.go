package main

import "fmt"

func main() {
	// Below approach is taken if we want to figure out later what we want to add to it
	// var colors map[string]string

	// Pretty much equivalent to the code above
	// colors := make(map[string]string)

	// Can add something to the map with the code below
	// colors["white"] = "#ffffff"


	colors := map[string]string{
		"red": "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}


	// Built in delete function to remove something from a map
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}