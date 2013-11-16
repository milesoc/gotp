package gotp

import "log"

type DevNullActor struct {
	GoActor
}

func (d *DevNullActor) Receive(msg Message) error {
	log.Println("received", msg)
	return nil
}