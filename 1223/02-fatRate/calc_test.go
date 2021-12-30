package main

import (
	"homework/1223/02-fatRate/BMRService/BMR"
	"testing"
)

func TestCalc_GetCompleteBMR(t *testing.T) {
	type args struct {
		b *BMR.BMR
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "calc a person",
			args: args{
				b: &BMR.BMR{
					Person: &BMR.Person{
						Name: "测试者",
						Height: 180,
						Weight: 80,
						Age: 25,
						Gender: "male",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := BMR.Calc{}
			c.GetCompleteBMR(tt.args.b)
			var zero float64 = 0
			if tt.args.b.Person.Bmi == zero {
				t.Errorf("failed to cacl bmi")
			}
			if tt.args.b.Person.FatRate == zero {
				t.Errorf("failed to cacl fatrate")
			}
			if tt.args.b.Advice == "" {
				t.Errorf("failed to get advice")
			}
		})
	}
}
