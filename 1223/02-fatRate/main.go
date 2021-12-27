package main

import (
	"fmt"
	"homework/1223/02-fatRate/BMRService"
	"homework/1223/02-fatRate/BMRService/BMR"
)

func main() {
	/** 初始化一个服务 Service struct {list *list.List} */
	s := new(BMRService.Service).Init()

	receiver := new(Receiver)
	/** 模拟一个生产的过程 */
	for {
		p, err := receiver.receiveArgs()
		if err != nil {
			fmt.Println(err.Error(), ",请重新输入")
			continue
		}
		/** 往服务里生产一个person数据 */
		s.Producer(p)
		/** 生产结束 */
		if receiver.IsBreak == true {
			fmt.Println("\n录入结束")
			break
		}
	}

	calc := new(BMR.Calc)
	/** 模拟一个生产的过程 */
	loop := 0
	for {
		loop++
		/** 消费person消息,生成一个单人BMR对象 */
		b := BMR.BMR{
			Person: s.Consumer(),
		}
		fmt.Printf("\n正在消费第%v个人\n", loop)
		/** 总的计算器 */
		calc.GetCompleteBMR(&b)
		fmt.Println("姓名:", b.Person.Name)
		fmt.Println("年龄:", b.Person.Age)
		fmt.Println("性别:", b.Person.Gender)
		fmt.Println("BMI:", b.Person.Bmi)
		fmt.Println("体脂率:", b.Person.FatRate)
		fmt.Println("建议:", b.Advice)
		/** 消费结束 */
		if s.IsNil() {
			fmt.Println("\n消费结束")
			break
		}
	}

}
