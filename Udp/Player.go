package Udp

import (
	"github.com/LoliE1ON/iron-go-server/Types"
	"net"
)

// Add new player to players slice
func AddNewPlayer(addr *net.UDPAddr, players *Types.UdpPlayers) {

	candidate := Types.UdpPlayer{
		IP:   addr.IP,
		Port: addr.Port,
	}

	if !Contains(players, candidate) {
		players.Push(candidate)
	}

}

// Search through each element in a slice
func Contains(slice *Types.UdpPlayers, item Types.UdpPlayer) bool {

	for _, s := range *slice {
		if s.IP.String() == item.IP.String() && s.Port == item.Port {
			return true
		}
	}

	return false
}
