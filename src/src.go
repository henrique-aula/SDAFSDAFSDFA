package src

import "fmt"

func Soma(a int, b int) int {
	if a >= 4 {
		return a - b
	}
	return a + b
}

func Print_soma(a int, b int) {
	fmt.Printf("%v\n", a+b)
}
