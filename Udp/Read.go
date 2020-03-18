package Udp

import (
	"encoding/json"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

var (
	Buffer []byte = make([]byte, 1024)
	Players Types.UdpPlayers
	Data Types.UdpResponse
)

func Read(connection *net.UDPConn) {

	for {

		// Read channel
		end, addr, err := connection.ReadFromUDP(Buffer)
		if err != nil {
			log.Fatalln("Error read from UDP", err)
			return
		}

		err = json.Unmarshal(Buffer[0:end], &Data)
		if err != nil {
			log.Println("Error Unmarshal TCP data", err)
			return
		}

		// Add new player
		AddNewPlayer(addr, &Players, connection, Data.Port)

		// Send Received data to all players
		for _, player := range Players {

			// Send
			_, err := player.Connection.WriteToUDP(Buffer[0:end], &net.UDPAddr{
				IP: player.IP,
				Port: player.Port,
			})

			if err != nil {
				log.Fatalln("Error write to UDP channel", err)
				return
			}

		}

	}

}