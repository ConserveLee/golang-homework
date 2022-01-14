package elevator

import (
	"fmt"
	"testing"
)

func Test_interface(t *testing.T) {
	var test PersonInterface = &XiaoMing{}
	s := test.Speak()
	fmt.Println(s)
}
