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
	done := make(chan struct{}) // 创建一个可以发送struct类型数据的channel
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 后台goroutine完成后向主goroutine发送完成信号
	}()                         // 另一个goroutine从连接中读取信息并打印服务器的响应
	mustCopy(conn, os.Stdin)    // main goroutine从标准输入流中读取内容并发送到服务器，关闭标准输入之后此函数返回继续往下执行
	conn.Close()
	<-done // 主goroutine接收到后台goroutine完成信号后退出程序
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
