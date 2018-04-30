package main

import (
	"net/http"
	"fmt"
	"log"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// http.HandlerFunc是类型转换而不是函数调用
	// db.list实现了handler类似的行为，但因为它没有方法，所以不满足http.Handler接口
	// 必须转换类型
	mux.HandleFunc("/list", db.list) // 简化http.HandlerFunc注册流程
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

/*
HandlerFunc是一个让函数值满足一个接口的适配器，
这里函数和这个接口仅有的方法有相同的函数签名。

package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
 */
