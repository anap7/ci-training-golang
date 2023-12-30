package main

import "fmt"

func main() {
	fmt.Println(Soma(10,10))
	fmt.Println(Soma(10,12))
}

func Soma(a int, b int) int {
	return a + b
}