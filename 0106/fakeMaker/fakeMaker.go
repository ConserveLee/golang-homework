package fakeMaker

import (
	"fmt"
	"homework/0106/building"
	"homework/0106/passenger"
	"math/rand"
)

type Maker struct {
	Building *building.Building
	Passengers *[]passenger.Passenger
}

// MakeBuilding
// 初始化一个建筑,此时按钮均未按
func (m *Maker) MakeBuilding() *building.Building {
	m.Building = building.Building{}.Init()
	return m.Building
}

func (m *Maker) MakeRandomPassengers() *[]passenger.Passenger {
	pNum := rand.Intn(19) + 1
	ps := make([]passenger.Passenger, pNum)
	for i := 0; i < pNum; i++ {
		ps = append(ps, m.MakeRandomPassenger(pNum, m.Building.TopFloor))
	}
	m.Passengers = &ps
	return m.Passengers
}

func (m *Maker) MakeRandomPassenger(pNum ,fNum int) passenger.Passenger {
	stayingFloor := rand.Intn(fNum - 1) + 1
	p := passenger.Passenger{
		Name: fmt.Sprintln("乘客", pNum),
		StayingFloor: stayingFloor,
	}
	p.ThinkToFloor(m.Building.TopFloor)
	return p
}