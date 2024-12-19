package solutions

import (
	"fmt"
	"time"
)

// concurrency Producer Consumer

type User struct {
	pause  chan bool
	resume chan bool
}

func (user *User) Producer() {
	for i := 0; i < 3; i++ { // Limit iterations for demonstration
		fmt.Println("Producer signaling pause")
		user.pause <- true
		time.Sleep(2 * time.Second)

		fmt.Println("Producer signaling resume")
		user.resume <- true
		time.Sleep(2 * time.Second)
	}
}

func (user *User) Consumer() {
	paused := false // Tracks whether the consumer is paused

	for {
		select {
		case <-user.pause:
			paused = true
			fmt.Println("Consumer received pause signal")
		case <-user.resume:
			paused = false
			fmt.Println("Consumer received resume signal")
		default:
			if !paused {
				// Perform work only if not paused
				fmt.Println("Consumer is doing something...")
			} else {
				fmt.Println("Consumer is paused...")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func ProducerConsumer() {
	fmt.Println("Starting program...")

	user := User{
		pause:  make(chan bool),
		resume: make(chan bool),
	}

	// Start Consumer in a goroutine
	go user.Consumer()

	// Start Producer
	user.Producer()

	fmt.Println("Program completed")
}
