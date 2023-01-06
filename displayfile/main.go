package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, _ := os.Open(os.Args[1])
		arr, _ := ioutil.ReadFile(os.Args[1])
		file.Read(arr)
		fmt.Printf(string(arr))
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
}
