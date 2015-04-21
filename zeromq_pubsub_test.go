package peer

import (
	"testing"

	"log"
	"time"
)

/*
	test a few subscribers and a publisher
*/
func TestPubSub(t *testing.T) {
	return // TEMP

	// create publisher
	msg := make(chan string)
	N := 3 // # subscribers

	// all quit channels
	quits := make([]chan bool, N+1)
	for i := 0; i < N+1; i++ { 
		quits[i] = make(chan bool)
	}

	
	// start publisher
	go publish("tcp://*:5556", msg, quits[0])

	// start subscribers
	for i := 1; i < N+1; i++ {
		go subscribe("tcp://localhost:5556", quits[i])
	}

	// wait a while then quit
	for i:=0;i<3;i++ {
		msg <- "moooo"
		time.Sleep(2*time.Second)
	}

	log.Println("Exiting...")
	for i := range quits {
		quits[i] <- true
	}
	
}

