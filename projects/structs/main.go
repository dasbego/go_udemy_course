package main

import "fmt"

// ContactInfo struct
type ContactInfo struct {
	email   string
	zipCode int
}

// Person struct
type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	me := Person{
		firstName: "Guillermo",
		lastName:  "Beltrán",
		ContactInfo: ContactInfo{
			email:   "gbego91@gmail.com",
			zipCode: 62900,
		},
	}
	pointToMe := *&me
	fmt.Println(pointToMe)
	pointToMe.updateName("Bego")
	me.print()
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
