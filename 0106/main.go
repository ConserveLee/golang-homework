package main

import (
	"fmt"
	"homework/0106/fakeMaker"
)

func main() {
	maker := fakeMaker.Maker{}
	maker.MakeBuilding()
	maker.MakeElevator()
	maker.MakeRandomPassengers()
	fmt.Println(maker)
	//摆烂了，现在只把建筑和乘客和电梯写出来了
}