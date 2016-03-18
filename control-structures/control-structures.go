package main

import "fmt"

func main() {

	if true == true {
		fmt.Println("true is true")
	} else {
		fmt.Println(" is false")
	}

	for i := 0; i < 20; i++ {
		fmt.Println("Go!")
	}

	for {
		fmt.Println("this will go forever")
		break
	}

	j := true
	for j {
		fmt.Println("while j is true")
		j = false
	}

	// Exercises:
	// 1. Print out the numbers 0-29 with 0-9 on line one, 10-19 on line two, and 20-29 on line three.
	for i := 0; i < 30; i++ {
		if i > 0 && i%10 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d ", i)
	}
}
