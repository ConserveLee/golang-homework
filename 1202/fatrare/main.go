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
	IsBreak  bool
	FatRate  float64
	FatLevel string
}

var (
	genderIndex = map[string]int{
		"男": 2,
		"女": 1,
	}
	fatRates struct {
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
	fmt.Printf("\n共录入%v人,%v人的平均体脂率值为:%v\n", fatRates.len, fatRates.len, fatRates.Sum/float64(fatRates.len))
}

func (h *human) readHumanBodyIndex(i int) float64 {
	fmt.Printf("\n正在录入第%v个人,\n请输入姓名:", i+1)
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
	h.FatRate = (1.2*h.Bmi + 0.23*(float64(h.Age)) - 5.4 - 10.8*float64(genderWeight)) / 100
	fmt.Printf("%s输入完成\n", h.Name)
	fmt.Printf("%s的BMI是:%v,体脂率是%v,%s\n", h.Name, h.Bmi, h.FatRate, h.getBodyTips())
	var whetherContinue string
	fmt.Print("是否录入下一个(y/n):")
	_, _ = fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		h.IsBreak = true
	}
	return h.FatRate
}

func (h *human) getBodyTips() string {
	var (
		thin = "体态:偏瘦。要多吃多锻炼,增强体质。"
		fit = "体态:标准。太棒了,要保持哦"
		overWeight = "体态:偏胖。吃完饭多散散步,消化消化"
		fat = "体态:肥胖。少吃点,多运动运动"
		veryFat = "体态:非常肥胖。健身游泳跑步,即刻开始"
		unknown = "我们不看未成年人的体脂率,变化太大,无法判断"
	)
	if h.Age < 18 {
		return unknown
	}
	if h.Gender == "男" {
		if h.Age >= 18 && h.Age <= 39 {
			if h.FatRate <= 0.1 {
				return thin
			} else if h.FatRate > 0.1 && h.FatRate <= 0.16 {
				return fit
			} else if h.FatRate > 0.16 && h.FatRate <= 0.21 {
				return overWeight
			} else if h.FatRate > 0.21 && h.FatRate <= 0.26 {
				return fat
			} else {
				return veryFat
			}
		} else if h.Age >= 40 && h.Age <= 59 {
			if h.FatRate <= 0.11 {
				return thin
			} else if h.FatRate > 0.11 && h.FatRate <= 0.17 {
				return fit
			} else if h.FatRate > 0.17 && h.FatRate <= 0.22 {
				return overWeight
			} else if h.FatRate > 0.22 && h.FatRate <= 0.27 {
				return fat
			} else {
				return veryFat
			}
		} else {
			if h.FatRate <= 0.13 {
				return thin
			} else if h.FatRate > 0.13 && h.FatRate <= 0.19 {
				return fit
			} else if h.FatRate > 0.19 && h.FatRate <= 0.24 {
				return overWeight
			} else if h.FatRate > 0.24 && h.FatRate <= 0.29 {
				return fat
			} else {
				return veryFat
			}
		}
	} else {
		if h.Age >= 18 && h.Age <= 39 {
			if h.FatRate <= 0.2 {
				return thin
			} else if h.FatRate > 0.2 && h.FatRate <= 0.27 {
				return fit
			} else if h.FatRate > 0.27 && h.FatRate <= 0.34 {
				return overWeight
			} else if h.FatRate > 0.34 && h.FatRate <= 0.39 {
				return fat
			} else {
				return veryFat
			}
		} else if h.Age >= 40 && h.Age <= 59 {
			if h.FatRate <= 0.21 {
				return thin
			} else if h.FatRate > 0.21 && h.FatRate <= 0.28 {
				return fit
			} else if h.FatRate > 0.28 && h.FatRate <= 0.35 {
				return overWeight
			} else if h.FatRate > 0.35 && h.FatRate <= 0.40 {
				return fat
			} else {
				return veryFat
			}
		} else {
			if h.FatRate <= 0.22 {
				return thin
			} else if h.FatRate > 0.22 && h.FatRate <= 0.29 {
				return fit
			} else if h.FatRate > 0.29 && h.FatRate <= 0.36 {
				return overWeight
			} else if h.FatRate > 0.36 && h.FatRate <= 0.41 {
				return fat
			} else {
				return veryFat
			}
		}
	}

}
