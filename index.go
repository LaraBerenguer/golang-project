package main

import "fmt"

func main() {
	number := 22
	if number >= 0 && number <= 9 {
		fmt.Println("The number has 1 digit")
	} else if number >= 10 && number <= 99 {
		fmt.Println("The number has 2 digits")
	} else if number >= 100 && number <= 999 {
		fmt.Println("The number has 3 digits")
	} else {
		fmt.Println("Number must be between 0 and 999")
	}
}
