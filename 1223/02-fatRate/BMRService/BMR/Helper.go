package BMR

func GetFakePerson() *Person {
	return &Person{
		Name:   "小明",
		Height: 170,
		Weight: 70,
		Age:    40,
		Gender: "male",
	}
}

func (b *BMR) GetCompleteCalc() {
	b.getFatRate()
	b.getAdvice()
}
