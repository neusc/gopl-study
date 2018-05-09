package main

import (
	"fmt"
	"log"
	"../../func/links"
	"os"
)

// 定义一个计数信号量，即容量有限的带缓存的channel
// 限制并发请求的数量
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 填充一个空的token
	list, err := links.Extract(url)
	<-tokens // 释放token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // 向worklist发送操作的数量

	// 此处将命令行参数传入worklist必须在单独的goroutine中进行
	// 否则会发生死锁
	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- { // 当没有运行中的crawl goroutine时推出循环
		list := <-worklist // 从channel接受数据
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
