package Udp

import (
	"fmt"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

var (
	Buffer []byte = make([]byte, 1024)
	Players Types.UdpPlayers
)

func Read(connection *net.UDPConn) {

	for {

		// Read channel
		end, addr, err := connection.ReadFromUDP(Buffer)
		if err != nil {
			log.Fatalln("Error read from UDP", err)
			return
		}

		// Add new player
		AddNewPlayer(addr, &Players)

		fmt.Println("Received ", string(Buffer[0:end]), " from ", addr)
		fmt.Println("Players ", Players)

		// Send Received data to all players
		for _, player := range Players {

			// Send
			_, err := connection.WriteToUDP(Buffer[0:end], &net.UDPAddr{
				IP: player.IP,
				Port: 5000,
			})
			if err != nil {
				log.Fatalln("Error write to UDP channel", err)
				return
			}

			fmt.Println("Send ", string(Buffer[0:end]), " to ", player.IP)
		}

	}

}