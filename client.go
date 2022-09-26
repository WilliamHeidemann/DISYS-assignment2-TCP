package main

import (
	"fmt"
	"time"
)

type client struct {
	tcpStream chan string
}
type clientPacket struct {
	message      string
	client       client
	index        int
	totalPackets int
}

func (client *client) connectToServer(server *server) {
	synPacket := clientPacket{
		"SYN",
		*client,
		1,
		1,
	}
	server.tcpStream <- synPacket
	response := <-client.tcpStream
	// error handling
	fmt.Printf("Client received the following packet: %s\n", response)
	time.Sleep(time.Second)
	if response == "SYN ACK" {
		synAckAckPacket := clientPacket{
			"SYN ACK ACK",
			*client,
			1,
			1,
		}
		server.tcpStream <- synAckAckPacket
		time.Sleep(time.Second)

		message := "Hello World!"
		var messagePackets []clientPacket
		messagePackets = client.marshalMessage(message)
		fmt.Println("Message packets created!")
		for _, packet := range messagePackets {
			if len(packet.message) == 0 {
				continue // Slice creates extra empty packets not meant to be sent
			}
			time.Sleep(time.Second)
			//fmt.Printf("Sending packet #%d containing \"%s\"\n", packet.index, packet.message)
			server.tcpStream <- packet
			//<-client.tcpStream // Packet was received
		}
	}
}

func (client *client) marshalMessage(message string) []clientPacket {
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
		}
		packets = append(packets, packet)
		fmt.Printf("Created packet %d out of %d\n", packet.index+1, packet.totalPackets)
		time.Sleep(time.Second)
		//fmt.Printf("Amount of packets created: %d\n", len(packets))
	}
	return packets
}
