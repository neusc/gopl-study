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
	//w := Wheel{Circle{Point{8, 8},5}, 20}}

	w := Wheel {
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n",w)
}

