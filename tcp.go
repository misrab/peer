
/*

	TCP async server and client

	TODO: 
		- tcp --> tls

*/

package peer 


import (
	"fmt"
	"time"

	"strconv"
	"net"
	"bufio"
)



func server(port int, messageChannel chan []byte) error {
	strport := strconv.Itoa(port)
	fmt.Println("Listening on tcp port "+ strport +"...")
	
	// listen on all interfaces
	ln, err := net.Listen("tcp", ":"+strport)
	if err != nil { return err }

	// accept connection on port
	conn, err := ln.Accept()
	if err != nil { return err }

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Print("Error reading from conn: ", err.Error())
			continue
		}

		// output message received
		fmt.Print("Message Received: ", string(message))

		messageChannel <- message

		conn.Write([]byte("received" + "\n"))
	}

	return nil
}

func client(host string, port int) error {
	strport := strconv.Itoa(port)
	address := host + ":" + strport

	// connect to this socket
	conn, err := net.Dial("tcp", address)
	if err != nil { return err }
	defer conn.Close()

	fmt.Println("Connected to tcp "+ address +"...")

	for {
		time.Sleep(time.Second)
		// read in input from stdin
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Text to send: ")
		// text, _ := reader.ReadString('\n')

		text := "what's your list"

		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Print("Message from server: "+ string(message))
	}

	return nil
}