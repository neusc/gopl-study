package main

import (
	"bytes"
	"fmt"
)

// 非负整数的slice集合
type IntSet struct {
	words []uint64
}

// has方法测试集合是否包含非负整数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Printf("%d => %d,%08b\n", x, word, bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// add方法添加非负整数x到集合
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 获取两个集合s和t的并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 打印集合作为字符串输出，格式'{1 2 3}'
// 依赖于接口和类型断言，fmt会直接调用用户自定义的String方法
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 求集合元素个数
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// 从集合删除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.words[word] != 0 && s.words[word]&(1<<bit) != 0 {
		s.words[word] &^= 1 << bit // 位操作符&^用于按位置零(AND NOT)
	}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1 9 144}
	fmt.Println(x.Len())
	x.Remove(7)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9 42}
	fmt.Println(y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String())           // {1 9 42 144}
	fmt.Println(&x)                   // {1 9 42 144}
	fmt.Println(x)                    // {[4398046511618 0 65536]} IntSet类型没有String方法，而*IntSet类型有String方法
	fmt.Println(x.Has(9), x.Has(123)) // true false
}
