package BMRService

import (
	"container/List"
	"homework/1223/02-fatRate/BMRService/BMR"
)

type Service struct {
	List *list.List
}

func (s *Service) Init() *Service {
	s.List = list.New()
	return s
}

func (s *Service) Producer(person *BMR.Person) {
	s.List.PushBack(person)
}

func (s *Service) Consumer() *BMR.Person {
	return s.List.Remove(s.List.Front()).(*BMR.Person)
}

func (s *Service) IsNil() bool {
	return s.List.Front() == nil
}
