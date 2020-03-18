package TcpEventListener

import (
	"fmt"
	"github.com/LoliE1ON/iron-go-server/Types"
)

// Send event
func SendEvent(channel chan Types.TcpEvent, event Types.TcpEvent) {

	channel <- event
	fmt.Println("Send event", event)

}