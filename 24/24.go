package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) calculateDistance(pointSecond *Point) float64 {
	return math.Sqrt(math.Pow(pointSecond.x-p.x, 2) + math.Pow(pointSecond.y-p.y, 2))
}

func main() {
	pointFirst := NewPoint(10.5, 12.25)
	pointSecond := NewPoint(2, 5.75)

	fmt.Println(pointFirst.calculateDistance(pointSecond) == pointSecond.calculateDistance(pointFirst))
}
