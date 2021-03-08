package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf257",
		"white": "#ffffff",
	}
	//Another ways to define a map

	//var colors map[string]string
	//colors := make(map[string]string)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is: ", hex)
	}
}
