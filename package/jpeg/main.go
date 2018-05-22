package main

import (
	"os"
	"fmt"
	"io"
	"image"
	"image/jpeg"
	_ "image/png"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

// 图像格式转换工具
func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in) // 从io.Reader接口读取数据并解码图像
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95}) //将image.Image类型的图像编码为jpeg格式
}
