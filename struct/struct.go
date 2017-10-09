package main

import (
	"fmt"
)

// Text struct
type Text struct {
	name string
}

type person struct {
	name string
	age  int
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

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}

	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51

	fmt.Println(sp.age)
}
