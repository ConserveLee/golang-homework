package main

import (
	"homework/1223/02-fatRate/BMRService"
	"homework/1223/02-fatRate/BMRService/BMR"
	"testing"
)

func TestService_ProducerAndConsumer(t *testing.T) {
	type fields struct {
		service *BMRService.Service
	}
	type args struct {
		person *BMR.Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "produce and consumer a person",
			fields: fields{
				service: new(BMRService.Service).Init(),
			},
			args:   args{
				person: &BMR.Person{
					Name: "测试者",
					Height: 180,
					Weight: 80,
					Age: 25,
					Gender: "male",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.service
			s.Producer(tt.args.person)
			s.Consumer()
		})
	}
}
