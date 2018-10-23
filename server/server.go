package main

//The server accepts messages from a number of local clients

import (
	"bufio"
	"fmt"
	"github.com/modmuss50/FlightServer/shared"
	"net"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	for {
		c, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(c.RemoteAddr().String() + " connected!")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			//Shall we send the error code back to the client?
			c.Close()
			fmt.Println(c.RemoteAddr().String() + " disconnected!")
			break
		}

		handleBaseStation(strings.TrimSpace(string(netData)), c)
	}
	c.Close()
}

func handleBaseStation(line string, c net.Conn) {
	//fmt.Println(line)

	basestation, err := shared.ParseBaseStation(line)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(basestation.MessageType)

}
