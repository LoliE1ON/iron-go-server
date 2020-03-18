package Tcp

import (
	"encoding/json"
	"github.com/LoliE1ON/iron-go-server/Tcp/TcpController"
	"github.com/LoliE1ON/iron-go-server/Tcp/TcpEventListener"
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

var (
	tcpConnection net.Conn
)

func Read(connection net.Conn) {

	var (
		tcpBuffer []byte = make([]byte, 1024)
		tcpData Types.TcpData
		tcpPlayers Types.TcpPlayers
	)

	tcpConnection = connection

	defer func() {
		event := Types.TcpEvent{
			Type:    "LeftPlayer",
			Payload: tcpData.Username,
		}
		TcpEventListener.SendEvent(Channel, event)
		TcpEventListener.ConnectedClients.RemoveByIp(connection.RemoteAddr().String())

		// Close the connection when you're done with it.
		err := connection.Close()

		log.Println("Close connection")
		if err != nil {
			log.Fatalln("Error close TCP connection", err)
			return
		}
	}()

	// Read the incoming connection into the buffer.
	for  {

		end, err := connection.Read(tcpBuffer)
		if err != nil {
			log.Println("Error reading TCP channel", err)
			return
		}

		err = json.Unmarshal(tcpBuffer[0:end], &tcpData)
		if err != nil {
			log.Println("Error Unmarshal TCP data", err)
			return
		}

		// Add new connected client
		client := TcpEventListener.Client{
			Ip: connection.RemoteAddr().String(),
			Connection: connection,
			Username:tcpData.Username,
		}
		TcpEventListener.ConnectedClients.Push(client)

		// Add new player
		AddNewPlayer(connection, tcpData.Username, &tcpPlayers)

		// Detect controller
		switch tcpData.Type {
		case "connect":
			TcpController.Connect(tcpConnection, tcpData, tcpPlayers, Channel)

		default:
			return
		}

		log.Println("Получено", tcpData)
		log.Println("Игроки", tcpPlayers)


	}


}
