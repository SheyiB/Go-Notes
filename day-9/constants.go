package main

import (
	"fmt"
	"math"
)

const color string = "blue"

func main(){
	fmt.Println(color)

	var name string
	fmt.Println(name)

	const fivesth = 5000000000

	const dilo = 3e20 / fivesth
	fmt.Println(dilo)

	fmt.Println(int64(dilo))


	fmt.Println(math.Sin(fivesth))

}