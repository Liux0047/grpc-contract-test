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
	msg    string
	weight float64
}

func main() {
	v := &Foo{
		name: "my name",
		age:  12,
		bar: &Bar{
			msg:    "hello",
			weight: 12.3,
		},
	}
	fmt.Println("type:", reflect.ValueOf(v).Elem().FieldByName("bar").Elem().FieldByName("weight").Type())

}
