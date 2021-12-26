package main

import "fmt"

type Advice struct {
}

var Advices = map[string]string {
	"thin": "体态:偏瘦。要多吃多锻炼,增强体质。",
	"fit": "体态:标准。太棒了,要保持哦",
	"overWeight": "体态:偏胖。吃完饭多散散步,消化消化",
	"fat": "体态:肥胖。少吃点,多运动运动",
	"veryFat": "体态:非常肥胖。健身游泳跑步,即刻开始",
	"unknown": "我们不看未成年人的体脂率,变化太大,无法判断",
}


func (Advice) getAdvice(p *Person, c *Calc)  {
	if p.Age < 18 || p.Age > 100 {
		p.Advice = Advices["unknown"]
		return
	}
	weight := int(p.FatRate*100)-c.getAgeWeight(p)
	fmt.Println(c.getAgeWeight(p), weight)
	if weight > 26 {
		p.Advice = Advices["veryFat"]
		return
	}
	if weight > 21 {
		p.Advice = Advices["fat"]
		return
	}
	if weight > 16 {
		p.Advice = Advices["overWeight"]
		return
	}
	if weight > 10 {
		p.Advice = Advices["fit"]
	}
	p.Advice = Advices["thin"]
	return
}

