package peer

import (
	"testing"

	"strconv"
)

func TestInitialise(t *testing.T) {
	println("Testing peer network...")

	N := 2
	group := make([]Peer, N)
	for index, _ := range group {
		group[index].Address = "tcp://localhost:555" + strconv.Itoa(index+1)
	}

	for index, g := range group {
		peers := append(group[:index], group[index+1:]...)
		Initialise(g.Address, peers)
	}

	//Initialise("tcp://*:5555", peers)
}