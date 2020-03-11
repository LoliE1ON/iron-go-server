package main

import (
	"github.com/LoliE1ON/iron-go-server/Tcp"
	"github.com/LoliE1ON/iron-go-server/Udp"
)

func main() {

	startUdp()
	startTcp()

}

func startUdp() {
	go Udp.Connect()
}

func startTcp() {
	Tcp.Connect()
}