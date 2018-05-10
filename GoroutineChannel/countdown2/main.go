package main

import (
	"os"
	"time"
	"fmt"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 监听键盘的return按键，从标准输入流中读入一个单字节
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown! Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// do something
		case <-abort:
			fmt.Println("Launch aborted.")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
