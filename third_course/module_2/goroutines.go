package main

import (
	"fmt"
	"sync"
)

/*
The race condition in the provided code occurs because both goroutines are incrementing the counter
variable concurrently without any synchronization.
The increment operation (counter++) is not atomic; it involves reading the current value,
incrementing it, and writing it back.
If both goroutines read the same value before either writes back,
one increment is lost, leading to an incorrect final value.
*/
func main() {
	var wg sync.WaitGroup
	counter := 0

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter--
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter) //Should be 0
}
