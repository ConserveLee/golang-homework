package main

import (
	"fmt"
)

type line struct {
	PointA point
	PointB point
}

type point struct {
	X float64
	Y float64
}

func main() {
	twoLine := [2]line{}
	for i, l := range twoLine {
		i = i+1
		fmt.Printf("\n正在输入输入线段%v\n:", i)
		fmt.Printf("请输入线段%v坐标A的X轴坐标:", i)
		_, _ = fmt.Scanln(&l.PointA.X)
		fmt.Printf("请输入线段%v坐标A的Y轴坐标:", i)
		_, _ = fmt.Scanln(&l.PointA.Y)
		fmt.Printf("请输入线段%v坐标B的X轴坐标:", i)
		_, _ = fmt.Scanln(&l.PointB.X)
		fmt.Printf("请输入线段%v坐标B的Y轴坐标:", i)
		_, _ = fmt.Scanln(&l.PointB.Y)
		twoLine[i-1] = l
	}
	if twoLine[0].isParallel(twoLine[1]) {
		fmt.Println("\n线段1和线段2平行")
		return
	}
	fmt.Println("\n线段1和线段2不平行")
}

func (l line) isParallel(l2 line) bool {
	return l != l2 && l2.slope() == l.slope()
}

func (l *line) slope() float64 {
	return (l.PointB.Y - l.PointA.Y)/(l.PointB.X - l.PointA.X)
}
