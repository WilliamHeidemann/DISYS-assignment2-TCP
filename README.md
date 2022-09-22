# DISYS-assignment2-TCP

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

Packets are a struct called clientPacket. It holds information about the message to be sent (string), the client it is sent from (client struct), 
the number of total packets and the index (int), meaning what number of packet it is itself (int).


b) Does your implementation use threads or processes? Why is it not realistic to use threads?

Our implementation use threads, though it is not realistic to use threads in real life as this communication would happen between different 
programs and not within the same program.


c) How do you handle message re-ordering?

As our implementation sends and receives one packet at a time, our implementation does not re-order after receiving the packets.


d) How do you handle message loss?

It doesn't:) 


e) Why is the 3-way handshake important?

It makes sure that both sides know that they are ready to transfer data, it is when you crossed the road as a child, look left, right and then left again:)
It is also necessary since both the server and the client needs to synchronize their segment sequence numbers.
