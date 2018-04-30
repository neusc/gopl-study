package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance   // 方法值，将方法绑定到特性接收器的函数
	fmt.Println(distanceFromP(q)) // "5"

	distance := Point.Distance  // 方法表达式，将第一个参数作为方法接收器
	fmt.Println(distance(p, q)) // "5"
}
