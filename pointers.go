package main

import "fmt"

func main() {
	num := 5
	double(num)
	fmt.Println(num)

	triple(&num)
	fmt.Println(num)

	a, b := 10, 20
	fmt.Printf("value of a(%d) and b(%d)\n", a, b)

	swap(&a, &b)

	fmt.Printf("value of a(%d) and b(%d)\n", a, b)

}

func double(num int)  {
	num *= 2
}

func triple(num *int) {
	*num *= 3
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
