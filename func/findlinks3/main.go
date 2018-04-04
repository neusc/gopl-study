package main

import (
	"fmt"
	"log"
	"os"
	"../links"
)

// worklist中的每一项调用f
// f返回的结果依次添加到worklist中
// 对于worklist的每一项最多调用依次f
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// 结构f返回的slice，一个个添加到worklist中
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
