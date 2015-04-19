package peer

import (
	"fmt"
    "log"

    zmq "github.com/pebbe/zmq4"
)


// host e.g. "tcp://*:5556"
func publish(host string, msg chan string, quit chan bool) {
	fmt.Println(host + " is starting to publish...")

	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PUB)
    defer context.Term()
    defer socket.Close()
    socket.Bind(host)

    for {
        select {
        case <- quit:
            log.Printf("Ending publish: %s...\n", host)
            return
        
        case m := <- msg:
            log.Printf("Publishing message: %s\n", m)
            socket.Send(m, 0)
        }

    	// do some fake "work"
        // time.Sleep(2*time.Second)

        // msg := fmt.Sprintf("%d %s", 1, "message")
        // socket.Send(msg, 0)
    }
}

func subscribe(host string, quit chan bool) {
    log.Printf("Subscribing to %s...\n", host)

	context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.SUB)
    defer context.Term()
    defer socket.Close()

    // var err error
    filter := ""


    ////  Subscribe to just one zipcode (whitefish MT 59937) //
    // fmt.Printf("Collecting updates from weather server for %s…\n", filter)
    socket.SetSubscribe(filter)
    socket.Connect(host) //"tcp://localhost:5556")

    for {
        select {
            case <- quit:
                log.Printf("Ending subscribe: %s...\n", host)
                return
            default:
                msg, _ := socket.Recv(0)
                log.Printf("Received from %s: %s\n", host, msg)
        }
    }
}