package Tcp

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "3001"
	TYPE = "tcp"
)

func Connect() {

	// Listen
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalln("Error listening", err)
		return
	}

	// Close listener
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Fatalln("Failed close TCP listener", err)
			return
		}
	}()

	fmt.Println("TCP: Listening on " + HOST + ":" + PORT)

	for {

		// Listen for an incoming connection.
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting", err)
			return
		}

		// Handle connections in a new goroutine.
		go Handle(connection)

	}
}

