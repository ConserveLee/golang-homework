package main

import (
	"fmt"
	bmi "github.com/armstrongli/go-bmi"//本地替换过的go-bmi
	"homework/1223/02-fatRate/BMRService/BMR"
)

func main() {
	/** 这里person用02的代码了,懒得写了 */
	b := BMR.BMR{
		Person: &BMR.Person{
			Name: "小明",
			Age: 30,
			Gender: "male",
			Height: 1.75,
			Weight: 75,
		},
	}
	err := *new(error)
	b.Person.Bmi, err = bmi.BMI(b.Person.Weight, b.Person.Height)
	if err != nil {
		fmt.Printf("get bmi error: %v", err)
		return
	}
	b.Person.FatRate = bmi.FatRate(b.Person.Bmi, b.Person.Age, b.Person.Gender)
	calc := BMR.Calc{}
	calc.GetAdvice(&b)
	fmt.Println("计算成功:\n姓名:", b.Person.Name)
	fmt.Println("年龄:", b.Person.Age)
	fmt.Println("性别:", b.Person.Gender)
	fmt.Println("BMI:", b.Person.Bmi)
	fmt.Println("体脂率:", b.Person.FatRate)
	fmt.Println("建议:", b.Advice)
}
