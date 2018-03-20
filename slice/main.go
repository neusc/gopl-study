package main

import "fmt"

func main() {
	// a=>array
	a := []int{0, 1, 2, 3, 4, 5}
	// reverse参数为slice
	reverse(a[:2])
	reverse(a[2:])
	reverse(a)
	fmt.Println(a)

	//长度为0的slice
	var s []int
	s = []int{}
	fmt.Println(s == nil)

	// append的使用
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// appentInt的使用
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appentInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	// nonempty的使用
	data := []string{"one","","three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	// remove
	p := []int{5,6,7,8,9}
	fmt.Println(remove(p, 2))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 向slice追加新元素
func appentInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 扩展slice
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap <= 2*len(x) {
			zcap = 2 * len(x)
		}
		// 空间不够，分配新数组
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// 去除slice中的空字符串
func nonempty(strings []string) []string  {
	i := 0
	for _,s := range strings {
		if s!= "" {
			 strings[i] = s
			 i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int  {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice) -1 ]
}


