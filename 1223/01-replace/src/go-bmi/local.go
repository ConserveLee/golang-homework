package gobmi

import "fmt"


func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

func FatRate(bmi float64, age int, gender string) float64 {
	var genderWeight float64 = 0
	if gender == "male" {
		genderWeight = 1
	}
	return (1.2*bmi + 0.23*(float64(age)) - 5.4 - 10.8*genderWeight) / 100
}