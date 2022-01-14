package elevator

type PersonInterface interface {
	Speak() string
}

type XiaoMing struct {
}

func (x XiaoMing) Speak() string {
	return "哈哈哈哈哈"
}
