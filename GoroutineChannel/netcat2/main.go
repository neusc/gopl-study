package main

import (
	"net"
	"log"
	"os"
	"io"
)

// 从连接中读取数据，将被读到的内容写到标准输出中
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn) // 另一个goroutine从连接中读取信息并打印服务器的响应
	mustCopy(conn, os.Stdin) // main goroutine从标准输入流中读取内容并发送到服务器
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
