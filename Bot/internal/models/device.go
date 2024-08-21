package model

import "fmt"

type Device struct {
	NwkAddr       string `json:"nwkAddr"`
	IeeeAddr      string `json:"ieeeAddr"`
	Flags         int32  `json:"flags"`
	Friendly_name string `json:"friendly_name"`
	Name          string
	Location      string
	Type          string
	State         St `json:"st"`
}

type St struct {
	State       string  `json:"state"`
	LastSeen    int32   `json:"last_seen"`
	Battery     int     `json:"battery"`
	Gas         string  `json:"gas"`
	Temperature float32 `json:"temperature"`
	Contact     bool    `json:"contact"`
}

func (device Device) String() string {
	return fmt.Sprintf(
		"Device:\n FriendlyName: *%s*\n NwkAddr: %s\n\n",
		device.Friendly_name,
		device.NwkAddr,
	)
}
