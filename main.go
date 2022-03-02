package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	name string
	age  int
	bar  *Bar
}

type Bar struct {
	msg string
}

func main() {
	v := &Foo{
		name: "my name",
		age:  12,
		bar: &Bar{
			msg: "hello",
		},
	}
	fmt.Println("type:", reflect.ValueOf(v).Elem().FieldByName("bar"))

}
