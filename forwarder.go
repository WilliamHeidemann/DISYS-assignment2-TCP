package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func forward(packet clientPacket) {
	var number = rand.Intn(10)
	switch number {
	case 0:
		fmt.Printf("DROPPING CLIENT PACKET %d\n", packet.index)
	default:
		send(packet)
	}
}

func forwardServerMessage(packet serverPacket) {
	packet.client.tcpStream <- packet.message
}

func send(packet clientPacket) {
	delay := rand.Intn(200) // wait for 0-2 seconds
	for i := 0; i < delay; i++ {
		time.Sleep(time.Millisecond)
	}
	packet.server.tcpStream <- packet
}

func sendACK(packet serverPacket) {
	var number = rand.Intn(10)
	switch number {
	case 0:
		index, _ := strconv.Atoi(packet.message)
		fmt.Printf("DROPPING ACK PACKET %d\n", index)
	default:
		index, _ := strconv.Atoi(packet.message)
		fmt.Printf("SENDING ACK INDEX IS %d\n", index)
		packet.client.ackStream <- index
	}
}
