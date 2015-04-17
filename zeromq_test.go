package peer

import (
	"testing"
)

func TestServer(t *testing.T) {
	go server()
	client()
}