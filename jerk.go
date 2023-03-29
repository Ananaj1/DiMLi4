package main

import (
	"fmt"
)

var a float64
var b float64
var d int
var c float64

func main() {
	fmt.Scan(&a)
	//fmt.Scan(&b)
	//fmt.Scan(&d)
	//fmt.Scan(&c)
	//fmt.Println(a != b)
	for fmt.Scan(&b); b <= a; b++ {
		fmt.Println("b now", b)
	}
}
