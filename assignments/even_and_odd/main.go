package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, element := range numbers {
		isEven := (element % 2) == 0
		if isEven == true {
			fmt.Println(element, "even")
		} else {
			fmt.Println(element, "odd")
		}
	}
}
