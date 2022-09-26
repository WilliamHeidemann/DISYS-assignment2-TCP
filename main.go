package main

import (
	"time"
)

func main() {
	server := &server{
		make(chan clientPacket),
	}

	go server.listenForPackets()

	client := &client{
		make(chan string),
	}

	go client.connectToServer(server)
	time.Sleep(time.Minute)
}
