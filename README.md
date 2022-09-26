# DISYS-assignment2-TCP

# Program description 

Our program starts by constructing a client and a server and communicates through a 
forwarder. This is implemented in golang using the channels feature.
First a three way hand shake takes place, to establish a connection between the server and the client.
When the connection has been established, the client marshalls its message into packets.
These packets are then sent concurrently to the forwarder, who is in charge of sending 
these messages to the server. 
A random slight delay will be added to each message during forwarding, so the packets do not
arrive in the correct order. 
The server is responsible for ordering its received messages.

When the original message has been ordered, it will be printed to the terminal and the 
program will terminate. 


a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

Packets are a struct called clientPacket. It holds information about the message to be sent (string), the client it is sent from (client struct), 
the number of total packets and the index (int), meaning what number of packet it is itself (int).


b) Does your implementation use threads or processes? Why is it not realistic to use threads?

Our implementation use threads, though it is not realistic to use threads in real life as this communication would happen between different 
programs and not within the same program.


c) How do you handle message re-ordering?

The server will receive all its packets in random order, but each packet contains a sequence number
called "index", which it can sort by. The server uses an inline anonymous function to sort its packets.

d) How do you handle message loss?

The forwarder can accidentally drop a packet, so it is never sent to the server.
To prevent information loss, the server will reply with an ACK-message to the client, after
receiving a packet. That way, if the client never receives an ACK for a packet, it will simply
try and resend the packet.


e) Why is the 3-way handshake important?

It makes sure that both sides know that they are ready to transfer data, it is like when you crossed the road as a child, look left, right and then left again:)
In real TCP, the client and the server need to synchronize their sequence numbers. This has 
been circumvented in our solution, by using a "total packets" number in each packet.