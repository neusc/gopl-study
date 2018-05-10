package main

import (
	"flag"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func main() {
	// 初始化目录为当前目录
	flag.Parse()
	roots := flag.Args() // 解析通过命令行参数传入的参数
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name()) // 获取子目录的路径
			walkDir(subdir, fileSizes)                 // 对子目录递归调用walkDir函数
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 返回一个目录下的入口列表slice
func dirents(dir string) []os.FileInfo {
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
