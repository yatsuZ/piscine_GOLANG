package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()
		} else {
			L := os.Args[1:]
			ord := false
			inser := false
			INSER := 0
			for i := range L {
				if len(L[i]) >= 9 {
					if L[i][:9] == "--insert=" && !(inser) {
						INSER = i
						L[i] = L[i][9:]
						inser = true
					}
				}
				if len(L[i]) >= 3 {
					if L[i][:3] == "-i=" && !(inser) {
						INSER = i
						L[i] = L[i][3:]
						inser = true
					}
				}
				if L[i] == "--order" && !(ord) || L[i] == "-o" && !(ord) {
					ord = true
				}
			}
			if inser && ord {
				ORD(INS(L, INSER))
			} else if inser {
				fmt.Println(INS(L, INSER))
			} else if ord {
				ORD(ROOM(L))
			} else {
				for j := range L {
					if L[j] != "--order" {
						fmt.Printf(L[j])
					}
				}
			}
		}
	} else {
		help()
	}
}

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func INS(L []string, Position int) string {
	dernier := L[Position]
	S := ""
	for i := range L {
		if i != Position && L[i] != "-o" && L[i] != "--order" {
			S = S + L[i]
		}
	}
	S = S + dernier
	return S
}

func ROOM(L []string) string {
	S := ""
	for i := range L {
		if L[i] != "--order" && L[i] != "-o" {
			S += L[i]
		}
	}
	return S
}

func ORD(s string) {
	var L []int
	for _, ajoute := range s {
		L = append(L, int(rune(ajoute)))
	}
	X := 0
	for decompte := range L {
		for i := range L {
			if L[decompte] > L[i] && decompte < i {
				X = L[decompte]
				L[decompte] = L[i]
				L[i] = X
			}
		}
	}
	for i := range L {
		z01.PrintRune(rune(L[i]))
	}
	z01.PrintRune(10)
}
