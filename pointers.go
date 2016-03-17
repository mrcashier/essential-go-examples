package main

import "fmt"

func main() {
	num := 5
	double(num)
	fmt.Println(num)

	triple(&num)
	fmt.Println(num)
}

func double(num int)  {
	num *= 2
}

func triple(num *int) {
	*num *= 3
}
