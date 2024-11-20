// //My program
//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
///*
//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”,
//respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
//*/
//
//func main() {
//	var name string
//	fmt.Print("Enter your name: ")
//	_, _ = fmt.Scan(&name)
//
//	var address string
//	fmt.Print("Enter your address: ")
//	_, _ = fmt.Scan(&address)
//
//	var personMap = map[string]string{
//		"name":    name,
//		"address": address,
//	}
//
//	marshal, err := json.Marshal(personMap)
//	if err != nil {
//		return
//	}
//
//	fmt.Println(string(marshal))
//}

// Other people's program

// The goal of this program is to practice working with RFCs in Go, using JSON as an example.

// This assignment is worth a total of 10 points:

// 3 points will be given if a program is written.
// 2 points will be given if the program compiles correctly.
// 5 points will be given if the program correctly prints a JSON object with keys ("name", "address") and they are associated with the name and address that was entered.

// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program,
// “makejson.go”.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	fmt.Printf("Enter your name:\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("Enter your address:\n")
	scanner.Scan()
	address = scanner.Text()

	m := map[string]string{"name": name, "address": address}

	barr, _ := json.Marshal(m)
	fmt.Printf("Your JSON:\n")
	fmt.Println(string(barr))

}
