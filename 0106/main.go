package main

import (
	"fmt"
	"homework/0106/fakeMaker"
	"math/rand"
	"time"
)

func main() {
	maker := fakeMaker.Maker{}
	maker.MakeBuilding()
	maker.MakeElevator()
	maker.MakeRandomPassengers()
	for _, passenger := range *maker.Passengers {
		fmt.Printf("%s在%v层\n",passenger.Name, passenger.StayingFloor)
	}
	ps := *maker.Passengers
	/** 模拟电梯运行 */
	for {
		/** 5s 随机取一名乘客 */
		passengerNum := rand.Intn(len(ps))
		passenger := ps[passengerNum]
		/** 按下乘客楼层的按钮 */
		passenger.PressButton(maker.Building)
		time.Sleep(5)
	}
	//摆烂了，现在只会写到这些，其他的后面补上
}