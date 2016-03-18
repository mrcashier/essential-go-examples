package main

import "fmt"

var (
	author  = "@mrcashier"
	Version = "0.0.1"
)

const (
	CCVisa       = "Visa"
	CCMastercard = "Mastercard"
)

func main() {
	fmt.Println("hello")
	fmt.Println("hello from fmt", "otro")

	var greeting string = "hello world"
	var a, b, c int = 1, 2, 3

	fmt.Println(greeting, a, b, c)

	var name = "Carlos Cajero"
	var d, e, f = 1, 2.0, "otro"

	fmt.Println(name, d, e, f)

	course := "Essential Go"
	x, y, z := 5, 6, 7

	fmt.Println(course, x, y, z)

	fmt.Printf("%d\n", x)

	fmt.Printf("%e\n", 10.500000000)

	fmt.Printf("%f\n", 10.500000000)
	fmt.Printf("%2.3f\n", 10.500000000) // 10.500

	fmt.Printf("%d %d", 20, 20)

}
