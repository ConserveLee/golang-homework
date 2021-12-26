package main

import "fmt"

func main() {
	p := getFakePerson()
	c := new(Calc)
	c.getFatRate(p)
	a := new(Advice)
	a.getAdvice(p, c)
	fmt.Println(p)
}

func getFakePerson() *Person {
	return &Person{
		Name: "小明",
		Height: 170,
		Weight: 83,
		Age: 40,
		Gender: "male",
	}
}