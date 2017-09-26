package main

import (
	"fmt"
)

// Text struct
type Text struct {
	name string
}

func (t *Text) setName(name string) {
	t.name = name
}

func (t Text) say() {
	fmt.Printf("Hello, %s\n", t.name)
}

func main() {
	Ming := Text{"ming ming"}
	XiaoXiao := Text{"Xiao Xiao"}

	Ming.setName("change name")
	Ming.say()

	XiaoXiao.say()
}
