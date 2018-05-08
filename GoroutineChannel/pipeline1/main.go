package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 计数器
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals) // 适当条件关闭channel，不再向其中发送数据
	}()

	// 接收naturals channel传来的整数并求平方，将结果传入squares channel
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // 退出循环当无法从channel获取到值，即channel已经关闭
			}
			squares <- x * x
		}
		close(squares)
	}()

	// 打印从squares channel接收到的数据
	for x := range squares {
		fmt.Println(x)
	}
}
