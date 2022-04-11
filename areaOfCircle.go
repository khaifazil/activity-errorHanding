package main

import (
	"errors"
	"fmt"
)

func calCircleArea(radius float64) (float64, error) {

	if radius <= 0 {
		return 0, errors.New("Radius must be positive")
	}

	const pi float64 = 3.14159
	area := pi * radius * radius
	return area, nil
}

func main() {

	if result, err := calCircleArea(-5); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
	fmt.Printf("Area of circle is: %v\n", result)
	}
}
