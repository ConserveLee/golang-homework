package building

import (
	"reflect"
	"testing"
)

func Test_building_GenBuilding(t *testing.T) {
	type TestFloor struct {
		Floors   map[Floor]bool
		TopFloor int
	}
	tests := []struct {
		name   string
		fields TestFloor
		want   *Building
	}{
		{
			name:   "能否生成楼层",
			fields: TestFloor{},
			want:   &Building{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Building{
				Floors:   tt.fields.Floors,
				TopFloor: tt.fields.TopFloor,
			}
			if got := b.GenBuilding(); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GenBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

