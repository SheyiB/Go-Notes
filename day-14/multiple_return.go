package main

import "fmt"

func vals() (int, int, string) {
	return 12, 20, "Hello"
}

func main() {

	a, b, e := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(e)

	_, _, c := vals()
	fmt.Println(c)
}