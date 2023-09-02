package main

import (
	"fmt"
)

var a int
var b int
var d int
var c float64

func main() {
	fmt.Scan(&a)
	for fmt.Scan(&b, &a); b >= a; a++ {
		fmt.Println("a now", a)
	}
}
