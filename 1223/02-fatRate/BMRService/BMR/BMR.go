package BMR

type BMR struct {
	Person *Person
	Advice string
}

var Advices = map[string]string{
	"thin":       "体态:偏瘦。要多吃多锻炼,增强体质。",
	"fit":        "体态:标准。太棒了,要保持哦",
	"overWeight": "体态:偏胖。吃完饭多散散步,消化消化",
	"fat":        "体态:肥胖。少吃点,多运动运动",
	"veryFat":    "体态:非常肥胖。健身游泳跑步,即刻开始",
	"unknown":    "年龄太小/太大,无法判断",
	"illegal":    "不正确的年龄",
}
