// A basic HTTP staticServer.
// By default, it serves the current working directory on port 8080.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type staticConfiguration struct {
	StaticPort string
	StaticPath string
}

func main() {
	file, _ := os.Open("../constants/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := staticConfiguration{}
	decodeErr := decoder.Decode(&conf)
	if decodeErr != nil {
		fmt.Println("Error", decodeErr)
	}
	log.Printf("listening on %s...", conf.StaticPort)
	err := http.ListenAndServe(conf.StaticPort, http.FileServer(http.Dir(conf.StaticPath)))
	log.Fatalln(err)
}
