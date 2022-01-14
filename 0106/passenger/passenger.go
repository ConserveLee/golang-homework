package passenger

import (
	"fmt"
	"homework/0106/building"
	"math/rand"
)

type Passenger struct {
	Name         string
	StayingFloor int
}

func (p *Passenger) ThinkToFloor(maxFloor int) int {
	wantFloor := rand.Intn(maxFloor - 1) + 1
	if wantFloor != p.StayingFloor {
		fmt.Printf("%s想去的楼层是：%v\n", p.Name, wantFloor)
		return wantFloor
	}
	return p.ThinkToFloor(maxFloor)
}

// PressButton
// 楼层的值为true,则被按过
func (p *Passenger) PressButton(f *building.Building) (ok bool) {
	pFloor := building.Floor(p.StayingFloor)
	if _, ok := f.Floors[pFloor]; ok {
		fmt.Printf("乘客%s按下了%v层的按钮", p.Name, pFloor)
		f.Floors[pFloor] = ok
	}
	return
}
