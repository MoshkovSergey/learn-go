package main

import (
	"fmt"
	"reflect"
)

// printTypeAndValue prints the type and value of the given input.
func printTypeAndValue(input interface{}) {
	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(input), input)
}

func main() {
	printTypeAndValue(25)      // Type: int, Value: 25
	printTypeAndValue(3.14)    // Type: float64, Value: 3.14
	printTypeAndValue("hello") // Type: string, Value: hello
	printTypeAndValue(true)    // Type: bool, Value: true
	printTypeAndValue('a')     // Type: int32, Value: 97
	printTypeAndValue(3 + 4i)  // Type: complex128, Value: (3+4i)
}
