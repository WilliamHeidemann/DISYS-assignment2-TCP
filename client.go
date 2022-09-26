package main

import (
	"fmt"
	"time"
)

type client struct {
	tcpStream chan string
	ackStream chan int
	acks      []int
}
type clientPacket struct {
	message      string
	client       client
	index        int
	totalPackets int
	server       server
}

func (client *client) connectToServer(server *server) {
	synPacket := clientPacket{
		"SYN",
		*client,
		1,
		1,
		*server,
	}
	server.tcpStream <- synPacket
	response := <-client.tcpStream
	// error handling
	fmt.Printf("Client received the following packet: %s\n", response)
	if response == "SYN ACK" {
		synAckAckPacket := clientPacket{
			"SYN ACK ACK",
			*client,
			1,
			1,
			*server,
		}
		forward(synAckAckPacket)

		message := "Hello World!"
		var messagePackets []clientPacket
		messagePackets = client.marshalMessage(message, server)
		fmt.Println("Message packets created!")

		for _, packet := range messagePackets {
			if len(packet.message) == 0 {
				continue // Slice creates extra empty packets not meant to be sent
			}
			//fmt.Printf("Sending packet #%d containing \"%s\"\n", packet.index, packet.message)
			go client.receiveAcks()
			go client.sendAndWaitForAck(packet)
			//<-client.tcpStream // Packet was received
		}
	}
}

func (client *client) sendAndWaitForAck(packet clientPacket) {
out:
	for {
		go forward(packet)
		time.Sleep(time.Second * 5)
		for index := range client.acks {
			if index == packet.index {
				break out
			}
		}
	}
}

func (client *client) receiveAcks() {
	for {
		ackIndex := <-client.ackStream
		client.acks = append(client.acks, ackIndex)
	}
}

func (client *client) marshalMessage(message string, server *server) []clientPacket {
	fmt.Println("Marshalling message...")

	packetSize := 5
	length := len(message) / packetSize
	if len(message)%packetSize != 0 {
		length += 1
	}
	fmt.Printf("Amount of packets to send: %d\n", length)

	var packets = make([]clientPacket, length)

	for i := 0; i < length; i++ {
		var sliceOfMessage string
		if i*packetSize+packetSize < len(message) {
			sliceOfMessage = message[i*packetSize : i*packetSize+packetSize]
		} else {
			sliceOfMessage = message[i*packetSize:]
		}
		packet := clientPacket{
			sliceOfMessage,
			*client,
			i,
			length,
			*server,
		}
		packets = append(packets, packet)
		fmt.Printf("Created packet %d out of %d\n", packet.index+1, packet.totalPackets)
	}
	return packets
}
