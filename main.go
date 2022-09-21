package main

type server struct {
	tcpStream chan clientPacket
}

type client struct {
	tcpStream chan string
}

type clientPacket struct {
	message string
	client  client
}

func main() {
	server := &server{
		make(chan clientPacket),
	}

	go server.listenForConnections()

	client := &client{
		make(chan string),
	}

	go client.connectToServer(server)
}

func (client *client) connectToServer(server *server) {
	synPacket := clientPacket{
		"SYN",
		*client,
	}
	server.tcpStream <- synPacket
	response := <-client.tcpStream
	// error handling
	if response == "SYN ACK" {
		synAckAckPacket := clientPacket{
			"SYN ACK ACK",
			*client,
		}
		server.tcpStream <- synAckAckPacket
	}
}

func (server *server) listenForConnections() {
	for {
		packet := <-server.tcpStream
		switch packet.message {
		case "SYN":
			packet.client.tcpStream <- "SYN ACK"
		case "SYN ACK ACK":
			var packets []clientPacket
			packets = make([]clientPacket, 10)
			for {
				packet = <-server.tcpStream
				// validation
				append(packets, packet)
			}
		}
	}
}
