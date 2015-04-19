package peer

import (
	"testing"

	"log"
	"time"
)



func TestInitialise(t *testing.T) {
	log.Println("Initialising peer...")

	// create publisher
	msg := make(chan string)
	quit1 := make(chan bool)
	go publish("tcp://*:5556", msg, quit1)

	
	// create subscribers
	quit2 := make(chan bool)
	go subscribe("tcp://localhost:5556", quit2)

	// wait a while then quit
	func() {
		for i:=0;i<10;i++ {
			msg <- "moooo"
			time.Sleep(2*time.Second)
		}
		
		quit2 <- true
		quit1 <- true
		log.Println("Exiting...")
	}()
}

