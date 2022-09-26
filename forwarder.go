package main

import (
	"fmt"
	"math/rand"
	"time"
)

func forward(packet clientPacket) {
	var number = rand.Intn(3)
	switch number {
	case 10:
		drop(packet)
	default:
		send(packet)
	}
}

func forwardServerMessage(packet serverPacket) {
	packet.client.tcpStream <- packet.message
}

func drop(packet clientPacket) {
	fmt.Printf("DROPPING PACKET %d\n", packet.index)
	//packet.server.tcpStream <- packet
}

func send(packet clientPacket) {
	delay := rand.Intn(200) // wait for 0-2 seconds
	for i := 0; i < delay; i++ {
		time.Sleep(time.Millisecond)
	}
	packet.server.tcpStream <- packet
}
