package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["carlos"] = 33
	ages["fer"] = 8
	ages["cris"] = 14

	fmt.Println(ages)
	fmt.Println(ages["carlos"])

	delete(ages, "carlos")

	fmt.Println(ages)

	ages2 := map[string]int{
		"carlos": 34,
		"cris":   14,
		"fer":    8,
	}

	fmt.Println(ages2)

	for n, a := range ages2 {
		fmt.Printf("%v is %v yeards old\n", n, a)
	}

	fmt.Println("Keys: ", getKeys(ages))
	fmt.Println("Keys: ", getKeys(ages2))

}

func getKeys(m map[string]int) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}
