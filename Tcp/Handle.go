package Tcp

import (
	"log"
	"net"
)

// Handles incoming requests.
func Handle(connection net.Conn) {
	
	// Read channel
	Read(connection)
	
	// Close the connection when you're done with it.
	err := connection.Close()
	if err != nil {
		log.Fatalln("Error close TCP connection", err)
		return
	}

}

