package main

import (
	"fmt"
)

// Test struct
type Test struct {
	name string
}

func (t *Test) setName(name string) {
	t.name = name
}

func (t Test) say() {
	fmt.Printf("Hello, %s\n", t.name)
}

func main() {
	Ming := Test{"ming ming"}
	XiaoXiao := Test{"Xiao Xiao"}

	Ming.setName("change name")
	Ming.say()

	XiaoXiao.say()
}
