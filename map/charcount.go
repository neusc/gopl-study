package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
	"strings"
)

func main() {
	counts := make(map[rune]int)  // unicode字符计数map
	var utflen [utf8.UTFMax + 1]int	// unicode码长度计数数组
	invalid := 0

	in := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	for {
		r,n,err	 := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n==1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c,n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i,n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
