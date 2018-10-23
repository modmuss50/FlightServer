package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func main() {
	conn, _ := net.Dial("tcp", "piplanes:30003")
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	defer conn.Close()
	for {
		message, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		handleBaseStation(message)
	}
}

func handleBaseStation(line string) {
	fmt.Println(line)
}
