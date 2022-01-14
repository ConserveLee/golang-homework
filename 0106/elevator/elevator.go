package elevator

import (
	"homework/0106/building"
	"homework/0106/passenger"
)

type Elevator struct {
	Floor     int
	MaxCarry  int
	Carrying  int
	Direction string
}

type ok bool

type EleInterface interface {
	Up() ok
	Down() ok
	CarryPassenger(p *passenger.Passenger) ok
	Running(b building.Building) *Elevator
}

func (e *Elevator) Up(b *building.Building) ok {
	if e.Floor < b.TopFloor {
		e.Floor++
		return true
	}
	return false
}

func (e *Elevator) Down() ok {
	if e.Floor > 0 {
		e.Floor--
		return true
	}
	return false
}

func (e *Elevator) CarryPassenger() ok {
	if e.Carrying == e.MaxCarry {
		return false
	}
	e.Carrying++
	return true
}
