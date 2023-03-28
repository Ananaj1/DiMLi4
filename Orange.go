package main

import "fmt"

var a int64
var b int64
var c int64

func main() {
	fmt.Println("введите а")
	fmt.Scan(&a)
	fmt.Println("введите b")
	fmt.Scan(&b)
	fmt.Println(a + b)

}
