package main

import (
	"fmt"
	"strings"
)

type LocomotionMethod int

const (
	Walk LocomotionMethod = iota
	Fly
	Slither
)

func (lm LocomotionMethod) String() string {
	return [...]string{"walk", "fly", "slither"}[lm]
}

type Animal struct {
	animalType       string
	eatenFood        string
	locomotionMethod LocomotionMethod
	spokenSound      string
}

func (animal Animal) Speak() string {
	return animal.spokenSound
}
func (animal Animal) Eat() string {
	return animal.eatenFood
}
func (animal Animal) Move() string {
	return animal.locomotionMethod.String()
}

func main() {
	cow := Animal{"cow", "grass", Walk, "moo"}
	bird := Animal{"bird", "worms", Fly, "peep"}
	snake := Animal{"snake", "mice", Slither, "hsss"}
	animalMap := map[string]Animal{"cow": cow, "bird": bird, "snake": snake}
	actionMap := map[string]func(Animal) string{"eat": Animal.Eat, "speak": Animal.Speak, "move": Animal.Move}

	var selectedAnimal string

	fmt.Println("Welcome to the farm! We have some animals for exhibition: a cow, a snake, and a bird")
	fmt.Print("Enter which animal you want to see: ")

	saveInput(&selectedAnimal)

	var selectedAction string

	fmt.Println("Enter what do you want to know about the animal between the following:")
	fmt.Println("Eat: what does it eat")
	fmt.Println("Move: how does it move")
	fmt.Println("Speak: what does it say")

	saveInput(&selectedAction)

	fmt.Println(actionMap[selectedAction](animalMap[selectedAnimal]))

}

func saveInput(value *string) {
	_, _ = fmt.Scan(value)
	*value = strings.ToLower(*value)
}
