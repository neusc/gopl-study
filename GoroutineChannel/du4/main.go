package main

import (
	"flag"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"time"
	"sync"
)

var done = make(chan struct{}) // 通过关闭channel向所有goroutine进行广播，便于取消后续操作

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// 初始化目录为当前目录
	flag.Parse()
	roots := flag.Args() // 解析通过命令行参数传入的参数
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 监测用户的取消操作
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) //对每一个walkDir的调用创建新的goroutine
	}

	go func() {
		n.Wait()
		close(fileSizes) // 计数器减为0是关闭channel
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// 释放fileSizes channel便于后续goroutine能够完成不阻塞
			for range fileSizes {
				// Do Something
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes channel被关闭，直接退出循环，带标签的break语句可以同时终结select和for两个循环
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done() // 结束一个walkDir的调用计数器减1
	if cancelled() {
		return // 每次调用walkDir都轮询是否触发了取消操作
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name()) // 获取子目录的路径
			go walkDir(subdir, n, fileSizes)           // 对子目录递归调用walkDir函数
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 自定义一个带缓存的channel，限制dirents函数的并发数量
var sema = make(chan struct{}, 20)

// 返回一个目录下的入口列表slice
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // 占用一个token
	case <-done:
		return nil // 取消对当前函数的调用
	}
	defer func() { <-sema }() // 释放一个token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
