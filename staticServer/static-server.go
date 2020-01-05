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

type StaticConfiguration struct {
	StaticServicePort string
	StaticPath        string
}

func main() {
	file, _ := os.Open("../constants/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := StaticConfiguration{}
	decodeErr := decoder.Decode(&conf)
	if decodeErr != nil {
		fmt.Println("Error", decodeErr)
	}
	log.Printf("listening on %s...", conf.StaticServicePort)
	err := http.ListenAndServe(conf.StaticServicePort, http.FileServer(http.Dir(conf.StaticPath)))
	log.Fatalln(err)
}
