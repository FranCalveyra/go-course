package main

import "fmt"

func main() {
	var input float64

	fmt.Print("Enter a floating point number: ")
	fmt.Scanf("%f", &input)

	// Truncate the floating point number by converting it to an integer
	truncated := int(input)

	fmt.Printf("The truncated integer is: %d\n", truncated)

}
