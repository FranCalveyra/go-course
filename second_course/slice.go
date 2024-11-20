package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer (or 'X' to exit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "X" || input == "x" {
			fmt.Println("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid integer or 'X' to exit.")
			continue
		}

		intSlice = append(intSlice, num)

		sort.Ints(intSlice)

		fmt.Println("Sorted slice:", intSlice)
	}
}

