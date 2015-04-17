package peer

import (
	"fmt"
    "time"

    zmq "github.com/pebbe/zmq4"
)


func publish(host string) {
	fmt.Println("Starting to publish...")

	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PUB)
    defer context.Term()
    defer socket.Close()
    socket.Bind(host) //"tcp://*:5556"

    // Wait for messages
    for {
    	// do some fake "work"
        time.Sleep(2*time.Second)

        msg := fmt.Sprintf("%d %s", 1, "im publishing something")
        socket.Send(msg, 0)
    }
}

func subscribe(host string) {
	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.SUB)
    defer context.Term()
    defer socket.Close()

    // var err error
    filter := ""


    ////  Subscribe to just one zipcode (whitefish MT 59937) //
    fmt.Printf("Collecting updates from weather server for %s…\n", filter)
    socket.SetSubscribe(filter)
    socket.Connect(host) //"tcp://localhost:5556")

    for i := 0; i < 101; i++ {
        msg, _ := socket.Recv(0)
        fmt.Printf("Received %s\n", msg)
        
    }
}