package Types

type UdpResponse struct {
	Type string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
	Position []float64 `json:"position,omitempty"`
	Rotation []float64 `json:"rotation,omitempty"`
	Port int `json:"port,omitempty"`
}
