package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	// map遍历顺序不固定
	for name,age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 顺序遍历map
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// 空的map
	var maps map[string]int
	fmt.Println(maps == nil)
	fmt.Println(len(maps) == 0)
	// maps["she"] = 25 // panic: assignment to entry in nil map

	// 判断0是真实存在map中的值还是不存在而返回的零值
	age, ok := ages["she"]
	if !ok {
		fmt.Printf("%d\t%t", age, ok)
	}
}

// 判断两个map是否含有相同的key和value
