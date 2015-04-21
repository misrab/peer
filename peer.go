package peer 

import (
	"log"
	"time"
)


var (
	DEFAULT_PEERS = []string{ "tcp://localhost:5000", }
)



// retrieve past peers based on database
func getPastPeers() []Peer {
	var result []Peer

	// try getting from database



	// if none return default
	if len(result) == 0 {
		log.Println("No past peers, getting from default peer...")

		result := make([]Peer, len(DEFAULT_PEERS))

		for i, address := range DEFAULT_PEERS {
			p := new(Peer)
			p.Address = address
			p.UpdatedAt = time.Now()
			result[i] = *p
		}
		
		return result
	}



	return result
}