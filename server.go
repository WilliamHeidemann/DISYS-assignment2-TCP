package main

import (
	"fmt"
	"time"
)

type server struct {
	tcpStream chan clientPacket
}

func (server *server) listenForPackets() {
	for {
		packet := <-server.tcpStream
		fmt.Printf("Server received the following packet: %s\n", packet.message)
		time.Sleep(time.Second)
		switch packet.message {
		case "SYN":
			packet.client.tcpStream <- "SYN ACK"
		case "SYN ACK ACK":
			fmt.Println("Connection between server and client established!")
			var packets []clientPacket
			packets = make([]clientPacket, 10)
			for {
				packet = <-server.tcpStream
				fmt.Printf("Server recieved packet #%d containing \"%s\"\n", packet.index, packet.message)
				// validation
				packets = append(packets, packet)
				//packet.client.tcpStream <- "Packet received!" // Den her linje virker ikke
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
