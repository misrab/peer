package peer

import (
	"fmt"
    "log"
    "time"

    zmq "github.com/pebbe/zmq4"
)

const (
    ZMQ_EAGAIN = 35 // nothing on socket, see http://api.zeromq.org/4-0:zmq-msg-recv#toc2
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
                log.Printf("Ending subscribe to %s...\n", host)
                return

            default:
                msg, err := socket.Recv(zmq.DONTWAIT)

                var num zmq.Errno
                if err != nil {
                    num = zmq.AsErrno(err)
                    // if nothing then wait
                    if num == ZMQ_EAGAIN {
                        time.Sleep(time.Second)
                        continue
                    } else { // else log and quit
                        log.Printf("Error, ending subscribe: %s\n", err.Error())
                        return
                    }
                }

                log.Printf("Received: %s\n", msg)
        }
    }
}