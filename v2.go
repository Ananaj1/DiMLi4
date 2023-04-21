package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(2)
}

func printCircleArea(radius int) {
	CircleAria, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Printf("Площадь круга: %f32 см.кв.\n\n", CircleAria)

}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус не может быть меньше нуля!")
	}
	return float32(radius) * float32(radius) * pi, nil

}
