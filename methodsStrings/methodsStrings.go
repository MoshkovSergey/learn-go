package main

import (
	"fmt"
	"strings"
)

var (
	name      string = "Sergey"
	lastName  string = "Moshkov"
	// age       int    = 30
	// isMarried bool   = false
)

func main() {
	compareStrings()
	containStrings()
	ReplaceStrings("Sergey")
}

// compareStrings compares the specified names by length and outputs the result
func compareStrings() {
	// Compare lastName with name
	fmt.Println(strings.Compare(lastName, name))
	
	// Compare name with "Sergey"
	fmt.Println(strings.Compare(name, "Sergey"))
	
	// Compare "Sergey" with name
	fmt.Println(strings.Compare("Sergey", name))
	
	// Compare name with itself
	fmt.Println(strings.Compare(name, name))
}

// containNames checks if the name contains certain letters
func containStrings()  {
	fmt.Println(strings.Contains(name, "e"))
	fmt.Println(strings.Contains(name, "y"))
	fmt.Println(strings.Contains(name, "z"))
}

func ReplaceStrings(name string)  {
	fmt.Println(strings.Replace(name, "S", "l", 1))
}
