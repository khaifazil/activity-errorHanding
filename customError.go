package main

import "fmt"

type invalidTriangleError struct {
	side1   float64
	side2   float64
	side3   float64
	message string
}

type invalidSideError struct {
	errSide float64
	message string
}

func (e *invalidTriangleError) Error() string {
	format := "%v, side 1: %v, side 2: %v, side 3: %v"

	return fmt.Sprintf(format, e.message, e.side1, e.side2, e.side3)
}

func (e *invalidSideError) Error() string {
	format := "%v, side: %v"

	return fmt.Sprintf(format, e.message, e.errSide)
}

func createTriangle(side1 float64, side2 float64, side3 float64) (float64, error) {

	invalidSideMsg := "Side is less than 0"
	invalidTrianglemsg := "Triangle cannot be formed with"

	//checks if each side is valid
	if side1 < 0  {
		return 0, &invalidSideError{side1, invalidSideMsg}
	} else if side2 < 0  {
		return 0, &invalidSideError{side2, invalidSideMsg}
	} else if side3 < 0  {
		return 0, &invalidSideError{side3, invalidSideMsg}
	}

	//checks to see if triangle can be formed
	if !((side1+side2) >side3 && (side3+side2) >side1 && (side1+side3) >side2) {
		return 0, &invalidTriangleError{side1, side2, side3, invalidTrianglemsg}
	}
		
	s := (side1+side2+side3)/2
	area := (s*(s-side1)*(s-side2)*(s-side3))/(s*(s-side1)*(s-side2)*(s-side3))
	return area, nil
}

func main() {
	s1 := -1.0
	s2 := 5.0
	s3 := 20.0

	area, err := createTriangle(s1, s2, s3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of triangle is formed is %.2f units squared.", area)
	}
}
