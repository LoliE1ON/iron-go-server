package Tcp

import (
	"fmt"
	"github.com/LoliE1ON/iron-go-server/Tcp/TcpEventListener"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

const (
	HOST = "5.180.138.37"
	PORT = "3001"
	TYPE = "tcp"
)

func Connect() {

	// Listen
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalln("Error listening TCP", err)
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

	// Create channel
	var channel chan Types.TcpEvent = make(chan Types.TcpEvent)

	for {

		// Listen for an incoming connection.
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting", err)
			return
		}

		// Event listener
		go TcpEventListener.EventListener(channel)

		// Handle connections in a new goroutine.
		go Handle(channel,connection)

	}
}

