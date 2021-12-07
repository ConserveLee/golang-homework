package main

import (
	"fmt"
	"math"
)

type human struct {
	Name     string
	High     float64
	Weight   float64
	Age      int
	Gender   string
	Bmi      float64
	IsBreak	 bool
	FatRare  float64
	FatLevel string
}

var (
	genderIndex = map[string]int{
		"男": 1,
		"女": 0,
	}
	fatRates struct{
		Sum float64
		len int
	}
)

func main() {
	var people [100]human
	fatRates.len = 1
	for i, person := range people {
		fatRates.Sum += person.readHumanBodyIndex(i)
		fatRates.len += i
		if person.IsBreak {
			break
		}
	}
	fmt.Printf("\n共有%v人,%v人的平均体脂率值为:%v\n", fatRates.len, fatRates.len, fatRates.Sum / float64(fatRates.len))
}

func (h *human) readHumanBodyIndex(i int) float64 {
	fmt.Printf("\n正在输入第%v个人,\n请输入姓名:", i+1)
	_, _ = fmt.Scanln(&h.Name)
	fmt.Print("请输入身高(cm):")
	_, _ = fmt.Scanln(&h.High)
	fmt.Print("请输入体重(kg):")
	_, _ = fmt.Scanln(&h.Weight)
	h.Bmi = h.Weight / math.Pow(h.High/100, 2)
	fmt.Print("请输入年龄:")
	_, _ = fmt.Scanln(&h.Age)
	fmt.Print("请输入性别(默认:男):")
	_, _ = fmt.Scanln(&h.Gender)
	genderWeight, ok := genderIndex[h.Gender]
	if !ok {
		fmt.Printf("性别不能为%s", h.Gender)
		h.Gender = "男"
		genderWeight = genderIndex[h.Gender]
	}
	h.FatRare = (1.2*h.Bmi + 0.23*(float64(h.Age)) - 5.4 - 10.8*float64(genderWeight)) / 100
	fmt.Printf("%s输入完成\n", h.Name)
	fmt.Printf("%s的BMI是:%v,体脂率是%v\n", h.Name, h.Bmi, h.FatRare)//todo 建议:体脂率对照表
	var whetherContinue string
	fmt.Print("是否录入下一个(y/n):")
	_, _ = fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		h.IsBreak = true
	}
	return h.FatRare
}
