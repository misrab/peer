package peer 


import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"time"
)


func server() {
	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.REP)
    defer context.Close()
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
        socket.Send([]byte(reply), 0)
    }
}


func client() {
	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.REQ)
    defer context.Close()
    defer socket.Close()

    fmt.Printf("Connecting to hello world server…")
    socket.Connect("tcp://localhost:5555")

    for i := 0; i < 10; i++ {
        // send hello
        msg := fmt.Sprintf("Hello %d", i)
        socket.Send([]byte(msg), 0)
        println("Sending ", msg)

        // Wait for reply:
        reply, _ := socket.Recv(0)
        println("Received ", string(reply))
    }
}