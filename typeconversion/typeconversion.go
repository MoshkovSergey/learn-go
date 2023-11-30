package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// printValues is a helper function to print the given values
func printValues(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	// Convert float64 to int
	cV1 := 1.9
	cV2 := int(cV1)
	printValues(cV1, cV2)

	// Convert int to float64
	cV3 := 1
	cV4 := float64(cV3)
	printValues(cV3, cV4)

	// Convert string to int
	cV5 := "50000000"
	cV6, err := strconv.Atoi(cV5)
	printValues(cV5, err, reflect.TypeOf(cV6))

	// Convert int to string
	cV7 := 50000000
	cV8 := strconv.Itoa(cV7)
	printValues(cV7, reflect.TypeOf(cV8))

	// Convert string to float64
	cV9 := "1.9"
	cV10, err := strconv.ParseFloat(cV9, 64)
	printValues(cV9, err, reflect.TypeOf(cV10))
}
