// Package rectangle include rectprops.go
package rectangle

import (
	"fmt"
	"math"
)

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

// Area function is calculating the area of rectangle
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal function
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
