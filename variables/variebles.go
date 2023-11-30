package main

import "fmt"

var pl = fmt.Println

func main() {
	var vName string = "Sergey"
	var v1, v2 int = 1, 2
	var v3 = "hello"
	v4 := 2.4
	v4 = 5.4
	pl(vName, v1, v2, v3, v4)
}