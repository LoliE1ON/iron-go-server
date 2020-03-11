package TcpController

import (
	"github.com/LoliE1ON/iron-go-server/Types"
	"log"
	"net"
)

func Connect(сonnection net.Conn, data Types.TcpData, players Types.TcpPlayers) {
	log.Println("connect to game")

	//сonnection.Write([]byte("Message received."))
}
