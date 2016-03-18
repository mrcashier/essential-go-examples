package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println("empty: ", nums)

	nums[4] = 100
	fmt.Println("set: ", nums)
	fmt.Println("get: ", nums[4])

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("slice ints: ", ints)

	ints = append(ints, 6, 7)
	fmt.Println("slice ints: ", ints)

	fmt.Println("0-2", ints[:2])
	fmt.Println("2-4", ints[2:4])
	fmt.Println("5-7", ints[4:])

	for j, val := range ints {
		fmt.Println(j, val)
	}

	fmt.Println("Sum: ", ints, sum(ints))

	ints2 := []int{7, 1, 6, 2, 5, 3, 4, 9, 8}
	fmt.Println("Sum: ", ints2, sum(ints2))
}

func sum(ints []int) int {
	sum := 0
	for _, e := range ints {
		sum += e
	}

	return sum
}
