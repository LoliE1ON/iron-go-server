package TcpEventListener

import (
	"encoding/json"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
)

// Listen events
func EventListener(channel chan Types.TcpEvent) {
	for {
		event := <- channel

		if event.Type == "ConnectionNewPlayer" {
			connectionNewPlayer(event)
		}

		if event.Type == "LeftPlayer" {
			leftPlayer(event)
		}

	}
}

func leftPlayer(event Types.TcpEvent) {

	for _, player := range ConnectedClients {

		// Preparing response
		response := Types.TcpResponseConnect{
			Type: "leftPlayer",
			Player: event.Payload.(string),
		}

		// Structure to Bytes
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("Marshaling output failed:", err)
			return
		}

		_, err = player.Connection.Write(jsonBytes)
		if err != nil {
			log.Println("TCP: Error write to TCP", err)
			return
		}

	}

}

func connectionNewPlayer(event Types.TcpEvent) {

	for _, player := range ConnectedClients {

		// Preparing response
		response := Types.TcpResponseConnect{
			Type: "connectNewPlayer",
			Player: event.Payload.(string),
		}

		// Structure to Bytes
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("Marshaling output failed:", err)
			return
		}

		_, err = player.Connection.Write(jsonBytes)
		if err != nil {
			log.Println("TCP: Error write to TCP", err)
			return
		}

	}

}
