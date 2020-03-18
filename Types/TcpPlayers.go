package Types

type TcpPlayer struct {
	IP string
	Username string
}

type TcpPlayers []TcpPlayer

// Push item
func (players *TcpPlayers) Push (player TcpPlayer) {
	*players = append(*players, player)
}
