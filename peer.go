package peer

import (
	"time"
)


const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type KnownNodes struct {
	Nodes []Node
	LastUpdated time.Time
}


type Node struct {
	Url string
	LastSeenActive time.Time
}


// main function to get into the network i.e. listen and respond
func JoinNetwork() {
	go server()
	time.Sleep(4)
	client()
}
