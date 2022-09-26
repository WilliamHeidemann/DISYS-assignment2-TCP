package main

import (
	"fmt"
	"sort"
)

type server struct {
	tcpStream chan clientPacket
}

type serverPacket struct {
	message string
	client  client
}

func (server *server) listenForPackets() {
	for {
		packet := <-server.tcpStream
		fmt.Printf("Server received the following packet: %s\n", packet.message)
		switch packet.message {
		case "SYN":
			synAckPacket := serverPacket{
				"SYN ACK",
				packet.client,
			}
			forwardServerMessage(synAckPacket)
		case "SYN ACK ACK":
			fmt.Println("Connection between server and client established!")

			var packets []clientPacket
			packets = make([]clientPacket, 0)

			for {
				packet := <-server.tcpStream
				fmt.Printf("Server recieved packet #%d containing \"%s\"\n", packet.index, packet.message)
				// validation
				packets = append(packets, packet)
				fmt.Printf("Packet length is: %d\n", len(packets))
				if len(packets) == packet.totalPackets {
					break
				}
			}

			// order packets
			sort.Slice(packets, func(i, j int) bool {
				return packets[i].index < packets[j].index
			})

			var originalMessage string
			for _, packet := range packets {
				originalMessage += packet.message
			}
			fmt.Println("Received message from client: \"" + originalMessage + "\"")
		}
	}
}
