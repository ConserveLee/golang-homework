package fakeMaker

import (
	"fmt"
	"homework/0106/building"
	"homework/0106/elevator"
	"homework/0106/passenger"
	"math/rand"
)

type Maker struct {
	Building *building.Building
	Passengers *map[int]passenger.Passenger
	Elevator *elevator.Elevator
}

// MakeBuilding
// 初始化一个建筑,此时按钮均未按
func (m *Maker) MakeBuilding() *building.Building {
	m.Building = building.Building{}.Init()
	return m.Building
}

func (m *Maker) MakeRandomPassengers() *map[int]passenger.Passenger {
	pNum := rand.Intn(19) + 1
	ps := make(map[int]passenger.Passenger, pNum)
	for i := 0; i < pNum; i++ {
		ps[i] = m.MakeRandomPassenger(i, m.Building.TopFloor)
	}
	m.Passengers = &ps
	return m.Passengers
}

func (m *Maker) MakeRandomPassenger(i ,fNum int) passenger.Passenger {
	stayingFloor := rand.Intn(fNum - 1) + 1
	return passenger.Passenger{
		Name: fmt.Sprintf("乘客%v", i+1),
		StayingFloor: stayingFloor,
	}
}

func (m *Maker) MakeElevator() *elevator.Elevator {
	m.Elevator = &elevator.Elevator{
		Floor:     1,
		MaxCarry:  20,
		Carrying:  0,
		Direction: "stop",
	}
	return m.Elevator
}