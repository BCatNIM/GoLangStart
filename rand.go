package main

import (
	"fmt"
)

var (
	year = 2005
)

func main() {

fmt.Println("THe year is 2100, should you leap?")

var leap =year%400 == 0 || (year%4 == 0 && year%100 != 0)

if leap {
	fmt.Println("Leap year - Loo before you leap!")
} else {
	fmt.Println("Short Feb - Keep your feet on the ground")
}

//	fmt.Println("You find yourself in a dimly lit cavern.")
//	var command = "forget that"
//	var exit = strings.Contains(command, "outside")
//	fmt.Println("You leave gthe cave:", exit)

//	var num = rand.Intn(10) +1
//	fmt.Println(num)

//	num = rand.Intn(10) +1
//	fmt.Println(num)
}