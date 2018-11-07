package main

import "fmt"
import "syscall/js"

func main() {
	fmt.Println("Hello, WebAssembly!")
}

func RegisterFunction(funcName string, myfunc func(i []js.value)) {

}
