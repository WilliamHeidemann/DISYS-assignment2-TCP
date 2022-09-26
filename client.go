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

		message := "We're soarin', flyin'\nThere's not a star in heaven\nThat we can't reach\nIf we're trying\nSo we're breaking free\nYou know the world can see us\nIn a way that's different than who we are\nCreating space between us\n'Til we're separate hearts\nBut your faith it gives me strength\nStrength to believe\nWe're breakin' free\n[Gabriella:]\nWe're soarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\n[Troy:]\nIf we're trying\n[Both:]\nYeah, we're breaking free\n[Troy:]\nOh, we're breakin' free\n[Gabriella:]\nOhhhh\n[Troy:]\nCan you feel it building\nLike a wave the ocean just can't control\n[Gabriella:]\nConnected by a feeling\nOhhh, in our very souls\n[Troy:]\nVery souls, ohhh\n[Both:]\nRising 'til it lifts us up\nSo everyone can see us\nWe're breakin' free\n[Gabriella:]\nWe're soarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\n[Troy:]\nIf we're trying\nYeah we're breaking free\n[Gabriella:]\nOhhhh runnin'\n[Troy:]\nClimbin'\nTo get to that place\n[Both:]\nTo be all that we can be\n[Troy:]\nNow's the time\n[Both:]\nSo we're breaking free\n[Troy:]\nWe're breaking free\n[Gabriella:]\nOhhh, yeah\nMore than hope\nMore than faith\n[Gabriella:]\nThis is true\nThis is fate\nAnd together\nWe see it comin'\n[Troy:]\nMore than you\nMore than me\nNot a want, but a need\n[Both:]\nBoth of us breakin' free\nSoarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\nIf we're trying\n[Troy:] Yeah we're breaking free\n[Gabriella:]\nBreaking free\nWere runnin'\n[Troy:]\nOhhhh, climbin'\n[Both:]\nTo get to the place\nTo be all that we can be\nNow's the time\n[Troy:] Now's the time\n[Gabriella:] So we're breaking free\n[Troy:] Ohhh, we're breaking free\n[Gabriella:] Ohhhh"
		var messagePackets []clientPacket
		messagePackets = client.marshalMessage(message, server)
		fmt.Println("Message packets created!")

		go client.receiveAcks()

		for _, packet := range messagePackets {
			if len(packet.message) == 0 {
				continue // Slice creates extra empty packets not meant to be sent
			}
			//fmt.Printf("Sending packet #%d containing \"%s\"\n", packet.index, packet.message)
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
		for i := 0; i < len(client.acks); i++ {
			if client.acks[i] == packet.index {
				fmt.Printf("ACK received! No longer sending packet %d\n", packet.index)
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
