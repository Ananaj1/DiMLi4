package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Х№р№ш№,благ№дарю за п№м№щь"
	b := strings.NewReplacer("№", "о")
	c := b.Replace(a)
	fmt.Println(c)
}
