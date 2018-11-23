package evend

import (
	"context"

	"github.com/temoto/vender/hardware/mdb"
)

type DeviceElevator struct {
	g DeviceGeneric
}

func (self *DeviceElevator) Init(ctx context.Context, mdber mdb.Mdber) error {
	// TODO read config
	self.g = DeviceGeneric{}
	return self.g.Init(ctx, mdber, 0xd0, "elevator")
}

func (self *DeviceElevator) ReadyChan() <-chan struct{} {
	return self.g.ready
}

func (self *DeviceElevator) CommandMove(position uint8) error {
	return self.g.CommandAction([]byte{0x03, position, 0x00})
}

// TODO poll, returns 0d
