package main

import (
	"os"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	// 该函数接收两个函数作为参数
	forEachNode(doc, startElement, endElement)

	return nil
}

// 递归遍历Html节点树，打印出页面结构
// 使用函数栈递归调用实现
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// %*s中的*会在字符串之前添加一些空格
		// 此处每次会先填充depth*2数量的空格，再输出""，最后再输出HTML标签
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

/*
<html>
  <head>
    <script>
    </script>
  </head>
  <body>
    <noscript>
    </noscript>
  </body>
</html>
 */
