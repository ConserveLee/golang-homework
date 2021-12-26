package main

type Person struct {
	Name     string
	Height   float64
	Weight   float64
	Age      int
	Gender   string
	Bmi      float64
	FatRate  float64
	Advice	 string
}

func (p *Person) getGenderWeight() float64 {
	if p.Gender == "male" {
		return 1
	}
	return 0
}