package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main () {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		if rand.Intn(10) == 0 {
			break
		}
	}
	if count == 0 {
		fmt. Println("liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}

}