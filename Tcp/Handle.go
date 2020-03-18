package Tcp

import (
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

var Channel chan Types.TcpEvent

// Handles incoming requests.
func Handle(channel chan Types.TcpEvent, connection net.Conn) {
	Channel = channel

	log.Println("New connection")

	// Read channel
	Read(connection)

}

