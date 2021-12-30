package BMR

import "math"

type Person struct {
	Name    string
	Height  float64
	Weight  float64
	Age     int
	Gender  string
	Bmi     float64
	FatRate float64
}

func (p *Person) getAgeWeight() int {
	if p.Gender == "female" {
		return p.getFemaleAgeWeight()
	}
	if p.Age > 60 {
		return 3
	}
	if p.Age > 39 {
		return 1
	}
	return 0
}

func (p *Person) getFemaleAgeWeight() int {
	if p.Age > 60 {
		return 12
	}
	if p.Age > 39 {
		return 11
	}
	return 10
}

func (p *Person) getGenderWeight() float64 {
	if p.Gender == "male" {
		return 1
	}
	return 0
}

func (p *Person) getBmi() {
	p.Bmi = p.Weight / math.Pow(p.Height/100, 2)
}

func (p *Person) getFatRate() {
	p.getBmi()
	p.FatRate = (1.2*p.Bmi + 0.23*(float64(p.Age)) - 5.4 - 10.8*p.getGenderWeight()) / 100
}
