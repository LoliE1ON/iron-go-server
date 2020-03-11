package Udp

import (
	"log"
	"net"
)

func Connect() {

	// Connect to UDP
	connection, err := net.ListenUDP("udp", &net.UDPAddr {
		IP: []byte{ 0, 0, 0, 0 },
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		log.Fatalln("Error connection to Udp", err)
		return
	}

	defer func() {
		err := connection.Close()
		if err != nil {
			log.Fatalln("Close connection error", err)
		}
	}()

	log.Println("Connected to Udp")

	// Read channel
	Read(connection)

}