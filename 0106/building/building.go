package building

import "math/rand"

type Floor int

type Building struct {
	Floors   map[Floor]bool
	TopFloor int
}


func (b Building) Init() *Building {
	b.TopFloor = rand.Intn(19) + 1
	b.Floors = make(map[Floor]bool, b.TopFloor)
	for i := 0; i < b.TopFloor; i++ {
		b.Floors[Floor(i)] = false
	}
	return &b
}