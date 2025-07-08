package main //main go package

import "fmt"

var g = "global"

func main() {
	l := "local"
	fmt.Println(l)
	fmt.Println(g)
}
