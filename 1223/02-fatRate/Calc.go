package main

import (
	"math"
)

type Calc struct {
}

func (Calc) getBmi(p *Person) {
	p.Bmi = p.Weight / math.Pow(p.Height/100, 2)
}

func (c Calc) getFatRate(p *Person) {
	c.getBmi(p)
	p.FatRate = (1.2*p.Bmi + 0.23*(float64(p.Age)) - 5.4 - 10.8*p.getGenderWeight()) / 100
}

func (c Calc) getAgeWeight(p *Person) int {
	if p.Gender == "female" {
		return c.getFemaleAgeWeight(p)
	}
	if p.Age > 60 {
		return 3
	}
	if p.Age > 39 {
		return 1
	}
	return 0
}

func (Calc) getFemaleAgeWeight(p *Person) int {
	if p.Age > 60 {
		return 12
	}
	if p.Age > 39 {
		return 11
	}
	return 10
}
