package Tcp

import (
	"encoding/json"
	"github.com/LoliE1ON/iron-go-server/Tcp/TcpController"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

var (
	Buffer []byte = make([]byte, 1024)
	Data Types.TcpData
	Players Types.TcpPlayers
	Connection net.Conn
)

func Read(connection net.Conn) {
	Connection = connection

	// Add new player
	AddNewPlayer(connection, &Players)

	// Read the incoming connection into the buffer.
	end, err := connection.Read(Buffer)
	if err != nil {
		log.Fatalln("Error reading TCP channel", err)
		return
	}

	err = json.Unmarshal(Buffer[0:end], &Data)
	if err != nil {
		log.Fatalln("Error Unmarshal TCP data", err)
		return
	}

	// Detect controller
	handleController(Data.Type)

	log.Println("Получено", Data)
	log.Println("Игроки", Players)

}

func handleController(controller string) {
	switch controller {
		case "connect":
			TcpController.Connect(Connection, Data, Players)

		default:
			return
	}
}
