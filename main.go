package main

import (
	"fmt"
)

func main() {
	server := &server{
		make(chan clientPacket),
		make(chan string),
	}

	go server.listenForPackets()

	client := &client{
		make(chan string),
		make(chan int),
		make([]int, 0),
	}

	go client.sendMessageTCP(server, "We're soarin', flyin'\nThere's not a star in heaven\nThat we can't reach\nIf we're trying\nSo we're breaking free\nYou know the world can see us\nIn a way that's different than who we are\nCreating space between us\n'Til we're separate hearts\nBut your faith it gives me strength\nStrength to believe\nWe're breakin' free\n[Gabriella:]\nWe're soarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\n[Troy:]\nIf we're trying\n[Both:]\nYeah, we're breaking free\n[Troy:]\nOh, we're breakin' free\n[Gabriella:]\nOhhhh\n[Troy:]\nCan you feel it building\nLike a wave the ocean just can't control\n[Gabriella:]\nConnected by a feeling\nOhhh, in our very souls\n[Troy:]\nVery souls, ohhh\n[Both:]\nRising 'til it lifts us up\nSo everyone can see us\nWe're breakin' free\n[Gabriella:]\nWe're soarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\n[Troy:]\nIf we're trying\nYeah we're breaking free\n[Gabriella:]\nOhhhh runnin'\n[Troy:]\nClimbin'\nTo get to that place\n[Both:]\nTo be all that we can be\n[Troy:]\nNow's the time\n[Both:]\nSo we're breaking free\n[Troy:]\nWe're breaking free\n[Gabriella:]\nOhhh, yeah\nMore than hope\nMore than faith\n[Gabriella:]\nThis is true\nThis is fate\nAnd together\nWe see it comin'\n[Troy:]\nMore than you\nMore than me\nNot a want, but a need\n[Both:]\nBoth of us breakin' free\nSoarin'\n[Troy:]\nFlyin'\n[Both:]\nThere's not a star in heaven\nThat we can't reach\nIf we're trying\n[Troy:] Yeah we're breaking free\n[Gabriella:]\nBreaking free\nWere runnin'\n[Troy:]\nOhhhh, climbin'\n[Both:]\nTo get to the place\nTo be all that we can be\nNow's the time\n[Troy:] Now's the time\n[Gabriella:] So we're breaking free\n[Troy:] Ohhh, we're breaking free\n[Gabriella:] Ohhhh")

	fmt.Println(<-server.messageToPrint)
}
