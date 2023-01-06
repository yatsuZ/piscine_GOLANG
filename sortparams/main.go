package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	machinbidule := []string(os.Args[1:])
	var L []int
	for c := range machinbidule {
		L = append(L, int(rune(machinbidule[c][0])))
	}
	//////////////je crée une liste d'int pour ensuite la rangé par orgre je prend uniquement le premier element
	x := ""
	X := 0
	for decompte := range L {
		for i := range L {
			if L[decompte] > L[i] && decompte < i {
				X = L[decompte]
				L[decompte] = L[i]
				L[i] = X
				x = machinbidule[decompte]
				machinbidule[decompte] = machinbidule[i]
				machinbidule[i] = x
			}
		}
	}
	/////////////////une fois que j'ai trié je l'affiche dans l'ordre :)
	for String := range machinbidule {
		for Rune := range machinbidule[String] {
			z01.PrintRune(rune(machinbidule[String][Rune]))
		}
		z01.PrintRune(10)
	}
}
