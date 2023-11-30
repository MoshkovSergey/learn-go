package main

import (
	"bufio"
	"log"
	"os"
)

// pl prints a string to the console
func pl(s string) {
	println(s)
}

// greetUser reads user input and greets the user
func greetUser() {
	pl("Hello, World!")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	pl("Hello, " + name)
}

func main() {
	greetUser()
}
