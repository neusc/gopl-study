package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	// 在Go的panic机制中，延迟函数的调用在释放堆栈信息之前
	// 所以可以在defer中打印函数发生异常时的堆栈报错信息
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
