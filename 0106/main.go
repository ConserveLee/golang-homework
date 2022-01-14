package main

import (
	"homework/0106/fakeMaker"
)

func main() {
	maker := fakeMaker.Maker{}
	maker.MakeBuilding()
	maker.MakeRandomPassengers()
	//摆烂了，现在只把建筑和乘客写出来了
}