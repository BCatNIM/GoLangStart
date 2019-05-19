package main

import (
    "fmt"
)

func main () {
	fmt. Println("There is a cavern entrance here and a path to the east.")
	var room = "cave"

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}