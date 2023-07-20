package main

import "fmt"

func main() {
	// different ways to create map
	// 1
	// var colors map[string]string

	// 2
	colors := make(map[string]string)

	//3
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// add key values to map
	colors["white"] = "#ffffff"
	colors["red"] = "#4bf745"
	colors["green"] = "#ffffff"

	// delete key value from map
	delete(colors, "white")

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
