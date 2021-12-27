package BMR

import "math"

type BMR struct {
	Person *Person
	Advice string
}

var Advices = map[string]string{
	"thin":       "体态:偏瘦。要多吃多锻炼,增强体质。",
	"fit":        "体态:标准。太棒了,要保持哦",
	"overWeight": "体态:偏胖。吃完饭多散散步,消化消化",
	"fat":        "体态:肥胖。少吃点,多运动运动",
	"veryFat":    "体态:非常肥胖。健身游泳跑步,即刻开始",
	"unknown":    "年龄太小/太大,无法判断",
	"illegal":    "不正确的年龄",
}

func (b *BMR) getAdvice() {
	if b.Person.Age < 0 || b.Person.Age > 150 {
		b.Advice = Advices["illegal"]
		return
	}
	if b.Person.Age < 18 || b.Person.Age > 100 {
		b.Advice = Advices["unknown"]
		return
	}
	weight := int(b.Person.FatRate*100) - b.Person.getAgeWeight()
	if weight > 26 {
		b.Advice = Advices["veryFat"]
		return
	}
	if weight > 21 {
		b.Advice = Advices["fat"]
		return
	}
	if weight > 16 {
		b.Advice = Advices["overWeight"]
		return
	}
	if weight > 10 {
		b.Advice = Advices["fit"]
	}
	b.Advice = Advices["thin"]
	return
}

func (b *BMR) getBmi() {
	b.Person.Bmi = b.Person.Weight / math.Pow(b.Person.Height/100, 2)
}

func (b *BMR) getFatRate() {
	b.getBmi()
	b.Person.FatRate = (1.2*b.Person.Bmi + 0.23*(float64(b.Person.Age)) - 5.4 - 10.8*b.Person.getGenderWeight()) / 100
}
