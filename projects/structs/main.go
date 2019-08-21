package main

import "fmt"

// Person struct
type Person struct {
	firstName string
	lastName  string
}

func main() {
	me := Person{
		firstName: "Guillermo",
		lastName:  "Beltrán",
	}
	me.firstName = "Bego"
	fmt.Println(me)
}
