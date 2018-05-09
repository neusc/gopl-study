package main

import (
	"fmt"
	"log"
	"../../func/links"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// 此处将命令行参数传入worklist必须在单独的goroutine中进行
	// 否则会发生死锁
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range worklist { // 从channel接受数据
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
