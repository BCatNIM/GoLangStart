package main

import (
    "fmt"
	"math/rand"
//	"time"
)

func main () {
	//here is my number
	var my_num = 9
	var rand_num = rand.Intn(10)
	var count = 0

	for my_num != rand_num {
		fmt.Print("My guess number:", rand_num, "--")
		fmt.Print("Your number was:", my_num, "\n")
		fmt.Println("")
		rand_num = rand.Intn(10)
		count++
		//time.Sleep(time.Second)
	}

	if my_num == rand_num {
		fmt.Print("Your number is ", my_num, ". It took me ", count, " guesses.\n")
	}

}