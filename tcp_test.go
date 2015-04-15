package peer 

import (
	"testing"
)

func TestServerClient(t *testing.T) {
	msgChan := make(chan []byte)
	go server(8081, msgChan)

	go client("127.0.0.1", 8081)

	for msg := range msgChan {
		t.Log(string(msg))
	}
}


// func TestClient(t *testing.T) {
// 	err := server(8081)
// 	if err != nil { t.Errorf("%v", err) }

// 	client("127.0.0.1", 8081)
// }