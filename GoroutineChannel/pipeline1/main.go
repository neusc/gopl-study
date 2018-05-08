package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 计数器
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// 接收naturals channel传来的整数并求平方，将结果传入squares channel
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 打印从squares channel接收到的数据
	for {
		fmt.Println(<-squares)
	}
}
