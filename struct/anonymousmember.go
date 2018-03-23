package main

import "fmt"

type Point struct {
	 X, Y int
}

// 结构体匿名成员Circle和Point都有自己的名字
// 就是命名的类型名字
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
}

