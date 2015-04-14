package peer 


import (
	"fmt"

	"net"
	"bufio"
	"os"
	"strings"
)

func server() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func client() {
	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil { panic(err) }
	defer conn.Close()

	for { 
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}