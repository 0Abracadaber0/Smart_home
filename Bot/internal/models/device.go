package model

type Device struct {
	NwkAddr       string `json:"nwkAddr"`
	IeeeAddr      string `json:"ieeeAddr"`
	Flags         int32  `json:"flags"`
	Friendly_name string `json:"friendly_name"`
	State         St     `json:"st"`
}

type St struct {
	State       string  `json:"state"`
	LastSeen    int32   `json:"last_seen"`
	Battery     int     `json:"battery"`
	Gas         string  `json:"gas"`
	Temperature float32 `json:"temperature"`
	Contact     string  `json:"contact"`
}
