package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	fmt.Println(p1.Distance(p2)) // 5

	p3 := NewPoint(-1, -1)
	p4 := NewPoint(2, 3)
	fmt.Println(p3.Distance(p4)) // 5

	p5 := NewPoint(5, 5)
	fmt.Println(p5.Distance(p5)) // 0
}
