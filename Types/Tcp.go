package Types

type TcpData struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Data string `json:"data"`
}

type TcpEvent struct {
	Type     string
	Payload interface{}
}

type TcpResponseConnect struct {
	Type string `json:"type"`
	Players []string `json:"players"`
	Player string `json:"player"`
}