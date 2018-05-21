package main

import "fmt"

// 通过显式的更改GOMAXPROCS的值，打印结果会改变
func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
