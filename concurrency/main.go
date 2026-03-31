package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true

}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal that the task is done
	close(doneChan)  // close the channel to signal that no more values will be sent
}

func main() {
	// dones := make([]chan bool, 4) // create a buffered channel to hold 4 done signals
	// done := make(chan bool)
	// dones[0] = make(chan bool)
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)
	for range done {
	}
}
