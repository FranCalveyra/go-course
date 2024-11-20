package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}
type Snake struct{}
type Bird struct{}

// Cow definitions

func (Cow) Eat() string   { return "grass" }
func (Cow) Move() string  { return "walk" }
func (Cow) Speak() string { return "moo" }

// Snake definitions

func (Snake) Eat() string   { return "mice" }
func (Snake) Move() string  { return "slither" }
func (Snake) Speak() string { return "hsss" }

// Bird definitions

func (Bird) Eat() string   { return "worms" }
func (Bird) Move() string  { return "fly" }
func (Bird) Speak() string { return "peep" }

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		cliEntries := strings.Fields(input)

		if len(cliEntries) != 3 {
			fmt.Println("Invalid command. Please use 'newanimal' or 'query' followed by the appropriate arguments.")
			continue
		}

		command, name, arg := cliEntries[0], cliEntries[1], cliEntries[2]

		switch command {
		case "newanimal":
			var animal Animal
			switch arg {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Invalid animal type. Please use 'cow', 'bird', or 'snake'.")
				continue
			}
			animals[name] = animal
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}
			switch arg {
			case "eat":
				fmt.Println(animal.Eat())
			case "move":
				fmt.Println(animal.Move())
			case "speak":
				fmt.Println(animal.Speak())
			default:
				fmt.Println("Invalid query. Please use 'eat', 'move', or 'speak'.")
			}

		default:
			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
		}
	}
}
