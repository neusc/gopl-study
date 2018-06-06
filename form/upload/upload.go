package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"io"
	"time"
	"crypto/md5"
	"html/template"
	"strconv"
)

func main() {
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 避免跨域问题
	fmt.Println("method:", r.Method)                   //获取请求的方法
	if r.Method == "GET" { // GET请求返回html页面
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("file_upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}