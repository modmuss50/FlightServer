package main

//The client connects to a locally running dump1090 and then pushes the messages to a remote server
//The idea of having this extra stage between dump1090 and the server is we can add extra data that dump1090 does not send, such as a username or email

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

var (
	RemoteServer net.Conn
)

func main() {

	err := connectToRemote()
	if err != nil {
		fmt.Println("Failed to connect to remote server!")
		fmt.Print(err)
		return
	}

	err = connectToDump1090()
	if err != nil {
		fmt.Println("Failed to connect to dump1090!")
		fmt.Print(err)
		return
	}
}

func connectToDump1090() error {
	conn, err := net.Dial("tcp", "piplanes:30003")
	if err != nil {
		return err
	}
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	defer conn.Close()
	for {
		message, err := tp.ReadLine()
		if err != nil {
			return err
		}
		handleBaseStation(message)
	}
}

func handleBaseStation(line string) {
	if len(line) == 0 {
		return //Skip empty lines
	}
	fmt.Println(line)
	sendMessage(line)
}

func connectToRemote() error {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		return err
	}
	RemoteServer = conn
	conn.Write([]byte("Hello, remote server, how are you doing today?\n"))
	return nil
}

func sendMessage(line string) {
	if RemoteServer == nil {
		fmt.Println("Not connected to remote!")
		return
	}
	fmt.Fprintf(RemoteServer, line+"\n")
}
