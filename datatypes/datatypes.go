package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf('a'))
	pl(reflect.TypeOf(3 + 4i))
}