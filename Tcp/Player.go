package Tcp

import (
	"github.com/LoliE1ON/iron-go-server/Types"
	"net"
)

// Add new player to players slice
func AddNewPlayer(connection net.Conn, players *Types.TcpPlayers) {

	candidate := Types.TcpPlayer{
		IP: connection.RemoteAddr().String(),
	}

	if !сontains(players, candidate) {
		players.Push(candidate)
	}

}

// Search through each element in a slice
func сontains(slice *Types.TcpPlayers, item Types.TcpPlayer) bool {

	for _, s := range *slice {
		if s.IP == item.IP {
			return true
		}
	}

	return false
}
