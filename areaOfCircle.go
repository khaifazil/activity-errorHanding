package main

import (
	"errors"
	"fmt"
	"math"
)

func calCircleArea(radius float64) (float64, error) {

	if radius <= 0 {
		return 0, errors.New("Radius must be positive")
	}

	area := math.Pi * radius * radius
	return area, nil
}

func main() {

	radius := 20.0
	if result, err := calCircleArea(radius); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Area of circle is: %v\n", result)
	}
}
