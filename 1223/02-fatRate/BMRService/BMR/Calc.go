package BMR

import "math"

type Calc struct {
}

func GetFakePerson() *Person {
	return &Person{
		Name:   "小明",
		Height: 170,
		Weight: 70,
		Age:    40,
		Gender: "male",
	}
}

func (c Calc) GetCompleteBMR(b *BMR) {
	c.getFatRate(b)
	c.getAdvice(b)
}

func (c Calc) getAdvice(b *BMR) {
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

func (c Calc) getFatRate(b *BMR) {
	b.Person.Bmi = b.Person.Weight / math.Pow(b.Person.Height/100, 2)
	b.Person.FatRate = (1.2*b.Person.Bmi + 0.23*(float64(b.Person.Age)) - 5.4 - 10.8*b.Person.getGenderWeight()) / 100
}
