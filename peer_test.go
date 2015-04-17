package peer

import (
	"testing"

	// "time"
	// "log"
	"strconv"
	"sync"
)

func TestInitialise(t *testing.T) {
	println("Testing peer network...")
	N := 3

	var wg sync.WaitGroup
	wg.Add(N)
	group := make([]Peer, N)
	for index, _ := range group {
		group[index].Address = "tcp://localhost:555" + strconv.Itoa(index+1)
	}

	// each address has the others as its peers
	for index, g := range group {
		var peers []Peer
		for i, _ := range group { if i != index { peers = append(peers, group[i]) } }
		Initialise(g.Address, peers, &wg)
	}


	wg.Wait()
}