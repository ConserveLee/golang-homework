package main

import (
	"fmt"
	"homework/1223/02-fatRate/BMRService"
	"homework/1223/02-fatRate/BMRService/BMR"
)

func main() {
	/** 初始化一个服务 Service struct {list *list.List} */
	s := new(BMRService.Service).Init()
	/** 往服务里生产一个person数据 */
	s.Producer(BMR.GetFakePerson())
	/** 消费person消息,生成一个单人BMR对象 */
	Bmr := BMR.BMR{
		Person: s.Consumer(),
	}
	calc := new(BMR.Calc)
	/** c.GetCompleteCalc(): c.getFatRate() & c.getAdvice() */
	calc.GetCompleteBMR(&Bmr)
	fmt.Println("person:", *Bmr.Person, "\nadvice:", Bmr.Advice)
}
