package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	r := []rune(s)
	return r[len(r)-1]
}

func main() {
	z01.PrintRune(FirstRune("Hello"))
	z01.PrintRune(FirstRune("Salut"))
	z01.PrintRune(FirstRune("Ola"))
	z01.PrintRune('\n')
}
