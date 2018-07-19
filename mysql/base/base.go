package base

import (
	"fmt"
)

type Base struct {
	str string
}

func (this *Base) Echo(str string) {
	fmt.Println(str)
}

func print(str string) {
	fmt.Println(str)
}
