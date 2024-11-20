package main

import (
	"fmt"
)

func main() {
	var numbers []int

	for len(numbers) != 10 {
		var input int
		_, _ = fmt.Scan(&input)
		numbers = append(numbers, input)
	}

	numbers = bubbleSort(numbers)

	fmt.Println("Sorted slice: ", numbers)

}

func bubbleSort(numbers []int) []int {
	for i := range numbers {
		for j := range numbers {
			if less(numbers[i], numbers[j]) {
				swap[int](&(numbers[i]), &(numbers[j]))
			}
		}
	}
	return numbers
}

func less(x int, y int) bool {
	return x < y
}

func swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}
