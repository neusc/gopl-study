package main

import (
	"time"
	"fmt"
)

// 主函数返回时，所有的goroutine都会被直接打断，然后退出程序
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)  // 菲波那切数列是耗时操作
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// 等待过程中展示动画小图标
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// 菲波那切数列递归算法
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
