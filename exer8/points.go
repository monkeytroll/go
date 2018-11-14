package exer8

import (
	"fmt"
	"math"
	"strings"
)

// ---------------- This is for testing ----------------
// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )

// func main() {
// 	a := Point{x: 1, y: 2}
// 	b := Point{3, 4.5}
// 	c := NewPoint(5, 6)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)

// 	fmt.Println(a.Norm())
// 	fmt.Println(b.Norm())
// 	fmt.Println(c.Norm())
// }
// ---------------- This is for testing ----------------

// TODO: The Point struct, NewPoint function, .String and .Norm methods

// Point is a pointer struct with an x and y coordinate
type Point struct {
	x float64
	y float64
}

// NewPoint creates a new Point struct
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// String provides the default print formatting for the Point struct
func (p Point) String() string {
	sX := strings.TrimRight(fmt.Sprintf("%f", p.x), "0")
	sX = strings.TrimRight(sX, ".")
	sY := strings.TrimRight(fmt.Sprintf("%f", p.y), "0")
	sY = strings.TrimRight(sY, ".")
	s := fmt.Sprintf("(%s, %s)", sX, sY)
	return s
}

// Norm calculates the Euclidean norm for a point
func (p Point) Norm() (norm float64) {
	norm = math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
	return
}
