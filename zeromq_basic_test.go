// Just a quick test to see if zeromq is configured properly,
// since it's only a binding to the C library

package peer

import (
	"testing"

	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func TestServer(t *testing.T) {
    // TEMP
    return

	println("\n============Testing ZeroMQ...==========\n")
	go server()

    // test multiple clients
	client()
}


func server() {
    fmt.Println("Starting server...")

	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.REP)
    defer context.Term()
    defer socket.Close()
    socket.Bind("tcp://*:5555")

    // Wait for messages
    for {
        msg, _ := socket.Recv(0)
        println("Received ", string(msg))

        // do some fake "work"
        time.Sleep(time.Second)

        // send reply back to client
        reply := fmt.Sprintf("World")
        socket.Send(reply, 0)
    }
}


func client() {
    fmt.Println("Starting client...")

	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.REQ)
    defer context.Term()
    defer socket.Close()

    fmt.Printf("Connecting to hello world server…")
    socket.Connect("tcp://localhost:5555")

    for i := 0; i < 5; i++ {
        // send hello
        msg := fmt.Sprintf("Hello %d", i)
        socket.Send(msg, 0)
        println("Sending ", msg)

        // Wait for reply:
        reply, _ := socket.Recv(0)
        println("Received ", string(reply))
    }
}