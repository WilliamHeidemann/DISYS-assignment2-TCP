package main

import (
	"fmt"
)

func main() {
	server := &server{
		make(chan clientPacket),
		make(chan string),
	}

	go server.listenForPackets()

	client := &client{
		make(chan string),
		make(chan int),
		make([]int, 0),
	}

	go client.connectToServer(server)

	fmt.Println(<-server.messageToPrint)
}
