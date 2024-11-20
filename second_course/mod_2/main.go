package main

import (
	"fmt"
	"math"
)

/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/

func main() {
	var a float64
	var v0 float64
	var s0 float64

	fmt.Print("Enter acceleration value: ")
	_, _ = fmt.Scanf("%f", &a)

	fmt.Print("Enter initial velocity value: ")
	_, _ = fmt.Scanf("%f", &v0)

	fmt.Print("Enter initial displacement value: ")
	_, _ = fmt.Scanf("%f", &s0)

	var t float64
	fmt.Print("Enter passed time: ")
	_, _ = fmt.Scanf("%f", &t)
	timeFunction := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Displacement: %f", timeFunction(t))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(t float64) float64 {
	// s = ½ a t2 + vot + so
	return func(t float64) float64 {
		return (0.5 * acceleration * math.Pow(t, 2)) + (initialVelocity * t) + initialDisplacement
	}
}
