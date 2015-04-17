// A node maintains a list of known peers. On startup, 
// it subscribes to them and listens

package peer

import (
	"log"
	"strconv"
	"sync"
)


// listens at given port e.g. "tcp://*:5555"
// subscribes to peers
func Initialise(host string, peers []Peer, wg *sync.WaitGroup) {
	go publish(host)

	//peers := getSavedPeers()
	for _, p := range peers {
		log.Println(host + " is subscribing to " + p.Address)
		go subscribe(p.Address)
	}

	

	println("intiialsed")
	defer wg.Done()
}




// dummy for now
func getSavedPeers() []Peer {
	N := 10
	result := make([]Peer, N)

	// dummy values
	for index, _ := range result {
		result[index].Address = "tcp://localhost:888" + string(strconv.Itoa(index+1))
	}

	return result
}