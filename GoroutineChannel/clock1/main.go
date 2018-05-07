package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main() {
	// listener对象监听网络端口上的连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept() // 阻塞直到一个新的连接被创建，返回一个net.Conn对象
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

// 处理一个完整的客户端连接，直到写入失败
// 最可能的原因是客户端主动断开连接，此时defer函数会关闭服务器上的该连接
// 然后返回主函数继续等待下一个连接请求
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("18:05:24\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
