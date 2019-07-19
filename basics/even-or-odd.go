package main

import "fmt"

func main() {

	fmt.Println("Hello, 世界")
	numbers := []int{1, 2, 3, 4, 5, 6}

	for _, number := range numbers {

		result := "odd"

		if number % 2 == 0 {
			result = "even"
		}

		fmt.Println(number, "is", result)
	}
}
