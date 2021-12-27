package BMRService

import (
	"container/list"
	"homework/1223/02-fatRate/BMRService/BMR"
)

type Service struct {
	list *list.List
}

func (s *Service) Init() *Service {
	s.list = list.New()
	return s
}

func (s *Service) Producer(person *BMR.Person) {
	s.list.PushBack(person)
}

func (s *Service) Consumer() *BMR.Person {
	return s.list.Remove(s.list.Front()).(*BMR.Person)
}
