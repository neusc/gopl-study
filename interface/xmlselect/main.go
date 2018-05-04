package main

import (
	"encoding/xml"
	"os"
	"io"
	"fmt"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // 存放结点元素名称的slice
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // 结点开始标签入栈
		case xml.EndElement:
			stack = stack[:len(stack)-1] // 结点结束标签出栈
		case xml.CharData:
			if containAll(stack, os.Args[1:]) { // 栈中有序包含通过命令行参数传入的元素名称
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// 判断在slice x中是否按顺序包含slice y中的元素
// [1,2,3,5] => [2,5]  true
func containAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
