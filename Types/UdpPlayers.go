package Types

import "net"

type UdpPlayer struct {
	IP   net.IP
	Port int
}

type UdpPlayers []UdpPlayer

// Push item
func (players *UdpPlayers) Push (player UdpPlayer) {
	*players = append(*players, player)
}