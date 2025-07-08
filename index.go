package main

import "fmt"

func main() {
	switch "Summer" {
	case "January", "March", "May", "July", "August", "October", "December":
		fmt.Println("This month has 31 days")
	case "April", "June", "September", "November":
		fmt.Println("This month has 30 days")
	case "February":
		fmt.Println("This month has 28 days")
	default:
		fmt.Println("Please, enter a valid month")
	}

	switch 2 {
	case 1:
		fmt.Println("This month has 31 days")
	case 2:
		fmt.Println("This month has 28 days")
	case 3:
		fmt.Println("This month has 31 days")
	case 4:
		fmt.Println("This month has 30 days")
	case 5:
		fmt.Println("This month has 31 days")
	case 6:
		fmt.Println("This month has 30 days")
	case 7:
		fmt.Println("This month has 31 days")
	case 8:
		fmt.Println("This month has 31 days")
	case 9:
		fmt.Println("This month has 30 days")
	case 10:
		fmt.Println("This month has 31 days")
	case 11:
		fmt.Println("This month has 30 days")
	case 12:
		fmt.Println("This month has 31 days")
	default:
		fmt.Println("Please, enter a valid month")
	}
}
