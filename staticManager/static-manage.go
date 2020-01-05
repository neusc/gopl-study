package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type StaticConfiguration struct {
	StaticPort string
	StaticPath string
	FilePath   string
}

type File struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"modTime"`
}

type ResponseData struct {
	StatusCode int64  `json:"statusCode"`
	Msg        string `json:"msg"`
	Data       []File `json:"data"`
}

type deleteParams struct {
	Name string `json:"name"`
}

var conf = StaticConfiguration{}

func main() {
	file, _ := os.Open("../constants/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Printf("listening on %s...", conf.StaticPort)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/static", getFileList)
	http.HandleFunc("/delete", deleteFile)
	log.Fatal(http.ListenAndServe(conf.StaticPort, nil))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 避免跨域问题
	fmt.Println("method:", r.Method)                   //获取请求的方法
	if r.Method == "POST" {                            // GET请求返回html页面
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(conf.StaticPath+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func getFileList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		fileInfo, err := ioutil.ReadDir(conf.StaticPath)
		if err != nil {
			fmt.Println("Read Dir error", err)
		}
		var fileList []File
		for _, file := range fileInfo {
			fileItem := File{Name: file.Name(), Path: conf.FilePath + file.Name(), Size: file.Size(), ModTime: file.ModTime().Unix()}
			fileList = append(fileList, fileItem)
		}
		w.WriteHeader(http.StatusOK)
		response := ResponseData{StatusCode: 0, Msg: "success", Data: fileList}
		// data, err := json.MarshalIndent(response, "", "     ")
		// fmt.Fprintf(w, "%s\n", data)
		json.NewEncoder(w).Encode(response)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	if r.Method == "POST" {
		var params deleteParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			panic(err)
		}
		deleteErr := os.Remove(conf.StaticPath + params.Name)
		if deleteErr != nil {
			fmt.Println("delete file err", deleteErr)
		}
		response := ResponseData{StatusCode: 0, Msg: "success", Data: nil}
		json.NewEncoder(w).Encode(response)
	}
}
