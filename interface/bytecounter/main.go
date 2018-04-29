package main

import "fmt"

type ByteCounter int

// 实现io Writer接口的任意类型
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) //将int类型转换为ByteCounter类型
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", len("hello")

	c = 0
	var name = "Dobby"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", len("hello, Dobby")
}
