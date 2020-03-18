package TcpEventListener

import (
	"log"
	"net"
)

type Client struct {
	Connection net.Conn
	Ip string
	Username string
}

type Clients []Client

var ConnectedClients Clients

// Push item
func (connectedClients *Clients) Push (client Client) {
	*connectedClients = append(*connectedClients, client)
}

// Push item
func (connectedClients *Clients) RemoveByIp (ip string) {
	var clients Clients

	for _, player := range *connectedClients {
		if player.Ip != ip {
			clients = append(clients, player)
		} else {
			log.Println("Remove IP ", ip)
		}
	}
	*connectedClients = clients
}