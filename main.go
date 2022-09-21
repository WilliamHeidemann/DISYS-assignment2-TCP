package main

import (
	"fmt"
	"time"
)

type server struct {
	tcpStream chan clientPacket
}

type client struct {
	tcpStream chan string
}

type clientPacket struct {
	message      string
	client       client
	index        int
	totalPackets int
}

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
	if response == "SYN ACK" {
		synAckAckPacket := clientPacket{
			"SYN ACK ACK",
			*client,
			1,
			1,
		}
		server.tcpStream <- synAckAckPacket
		message := "Hello World!"
		var messagePackets []clientPacket
		messagePackets = client.marshalMessage(message)
		fmt.Println("Message packets created!")
		for i, packet := range messagePackets {
			time.Sleep(time.Second)
			fmt.Printf("Sending packet #%d containing \"%s\"\n", i, packet.message)
			server.tcpStream <- packet
			//<-client.tcpStream // Packet was received
		}
	}
}

func (server *server) listenForPackets() {
	for {
		packet := <-server.tcpStream
		switch packet.message {
		case "SYN":
			packet.client.tcpStream <- "SYN ACK"
		case "SYN ACK ACK":
			fmt.Println("Connection between server and client established!")
			var packets []clientPacket
			packets = make([]clientPacket, 10)
			for {
				packet = <-server.tcpStream
				fmt.Printf("Packet #%d received\n", packet.index)
				// validation
				packets = append(packets, packet)
				//packet.client.tcpStream <- "Packet received!"
				//fmt.Println("Ack packet returned")
				if packet.index == packet.totalPackets-1 {
					fmt.Println("All packets were received!")
					break
				}
			}
			var originalMessage string
			for _, packet := range packets {
				originalMessage += packet.message
			}
			fmt.Println("Received message from client: \"" + originalMessage + "\"")
		}
	}
}

func (client *client) marshalMessage(message string) []clientPacket {
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
		fmt.Printf("Amount of packets created: %d\n", len(packets))
	}
	return packets
}
