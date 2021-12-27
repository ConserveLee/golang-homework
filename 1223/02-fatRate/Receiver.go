package main

import (
	"errors"
	"fmt"
	"homework/1223/02-fatRate/BMRService/BMR"
)

type Receiver struct {
	Loop    int
	IsBreak bool
}

var genders = map[string]bool{
	"male":   true,
	"female": true,
}

func (r *Receiver) receiveArgs() (*BMR.Person, error) {
	p := new(BMR.Person)
	fmt.Printf("\n正在录入第%v个人,\n请输入姓名:", r.Loop+1)
	_, _ = fmt.Scanln(&p.Name)
	fmt.Print("请输入身高(cm):")
	_, _ = fmt.Scanln(&p.Height)
	if p.Height <= 0 {
		return nil, errors.New("身高不能小于0")
	}
	fmt.Print("请输入体重(kg):")
	_, _ = fmt.Scanln(&p.Weight)
	if p.Weight <= 0 {
		return nil, errors.New("体重不能小于0")
	}
	fmt.Print("请输入年龄(age):")
	_, _ = fmt.Scanln(&p.Age)
	if p.Age < 0 || p.Age > 150 {
		return nil, errors.New("非法的年龄")
	}
	fmt.Print("请输入性别(male/female):")
	_, _ = fmt.Scanln(&p.Gender)
	_, ok := genders[p.Gender]
	if !ok {
		return nil, errors.New(fmt.Sprintf("性别不能为%s", p.Gender))
	}
	fmt.Printf("%s输入完成\n", p.Name)
	r.Loop++
	var whetherContinue string
	fmt.Print("是否录入下一个(y/n):")
	_, _ = fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		r.IsBreak = true
	}
	return p, nil
}
