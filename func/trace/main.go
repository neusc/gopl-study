package main

import (
	"time"
	"log"
)

func bigSlowOperation() {
	// 末尾的括号必须有，否则本该进入时执行的操作变成退出时执行
	// 本该退出时执行的操作永远不会被执行
	defer trace("bigSlowOperation")()
	// 模拟许多操作
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
