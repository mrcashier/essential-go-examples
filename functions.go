package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println(double(5))

	myName := "Carlos Cajero"
	first, _ := parseName(myName)

	fmt.Println(first)

	innerGreet := func(name string) {
		fmt.Println("Hello", name)
	}

	innerGreet(myName)

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)

	}
}

func double(n int) int {
	return n * 2
}

func parseName(name string) (first, last string) {
	log.Println("Received Params", first, last)
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
