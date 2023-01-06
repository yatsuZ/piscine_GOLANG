package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		if os.Args[0] == "echo" {
			for _, m := range os.Args[2] {
				z01.PrintRune(rune(m))
			}
		}
		for o, Z := range os.Args[1:] {
			if Z == "|" {
				for _, p := range os.Args[o-1] {
					z01.PrintRune(rune(p))
				}
			}
			file, err := os.Open(Z)
			if err == nil {
				arr, _ := ioutil.ReadFile(Z)
				file.Read(arr)
				s := string(arr)
				for _, v := range s {
					z01.PrintRune(rune(v))
				}
			} else {
				machin := "ERROR: open " + Z + ": no such file or directory\n"
				for _, l := range machin {
					z01.PrintRune(rune(l))
				}
				os.Exit(1)
			}
		}
	}
}
