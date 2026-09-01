package main

import "fmt"

func main() {
	var a [5]int
	var b [3]string
	var c [3]bool

	fmt.Println("array:", a)
	fmt.Println("array2", b)
	fmt.Println("array3", c)
	a[0] = 1
	fmt.Println("set:", a)

	fmt.Println(len(a))
}
