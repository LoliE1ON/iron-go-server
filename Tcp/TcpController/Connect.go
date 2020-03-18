package TcpController

import (
	"encoding/json"
	"github.com/LoliE1ON/iron-go-server/Tcp/TcpEventListener"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

func Connect(сonnection net.Conn, data Types.TcpData, players Types.TcpPlayers, channel chan Types.TcpEvent) {

	var usernames []string

	// Serialize all players
	for _, player := range TcpEventListener.ConnectedClients {
		if player.Username != data.Username {
			usernames = append(usernames, player.Username)
		}
	}

	// Preparing response
	response := Types.TcpResponseConnect{
		Type: "spawnOldPlayers",
		Players: usernames,
	}

	// Structure to Bytes
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Marshaling output failed:", err)
		return
	}

	log.Println("Connect to game")
	event := Types.TcpEvent{
		Type:    "ConnectionNewPlayer",
		Payload: data.Username,
	}

	TcpEventListener.SendEvent(channel, event)

	// Response
	_, err = сonnection.Write(jsonBytes)
	if err != nil {
		log.Println("TCP: Error write to TCP", err)
		return
	}

}
