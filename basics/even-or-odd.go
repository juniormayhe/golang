package main

import "fmt"

func main() {

	fmt.Println("Check if numbers are even or odd")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {

		result := "odd"

		if number % 2 == 0 {
			result = "even"
		}

		fmt.Println(number, "is", result)
	}
}
