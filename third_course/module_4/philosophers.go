package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const numTimesToEat = 3

type Philosopher struct {
	id        int
	leftChop  *sync.Mutex
	rightChop *sync.Mutex
}

func (p Philosopher) eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()
	for i := 0; i < numTimesToEat; i++ {
		host <- true // Request permission from the host
		p.leftChop.Lock()
		p.rightChop.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second) // Simulate eating
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightChop.Unlock()
		p.leftChop.Unlock()
		<-host                  // Release permission back to the host
		time.Sleep(time.Second) // Simulate thinking
	}
}

func main() {
	var wg sync.WaitGroup
	chopsticks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	philosophers := make([]Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{
			id:        i + 1,
			leftChop:  chopsticks[i],
			rightChop: chopsticks[(i+1)%numPhilosophers],
		}
	}

	host := make(chan bool, 2) // Host allows up to 2 philosophers to eat concurrently

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat(&wg, host)
	}

	wg.Wait()
}
