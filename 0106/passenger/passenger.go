package passenger

import (
	"homework/0106/building"
	"math/rand"
)

type Passenger struct {
	Name         string
	ToFloor      int
	StayingFloor int
}

func (p *Passenger) ThinkToFloor(maxFloor int) int {
	wantFloor := rand.Intn(maxFloor - 1) + 1
	if wantFloor != p.StayingFloor {
		return wantFloor
	}
	return p.ThinkToFloor(maxFloor)
}

// PressButton
// 楼层的值为true,则被按过
func (p *Passenger) PressButton(f *building.Building) (ok bool) {
	pFloor := building.Floor(p.StayingFloor)
	if _, ok := f.Floors[pFloor]; ok {
		f.Floors[pFloor] = ok
	}
	return
}
