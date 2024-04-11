package main

import "fmt"

func main(){
	for n:=0; n<100; n++{
		if n%5 == 0 && n %3 == 0{
			fmt.Println("Fizzbuzz")
		} else if n%5 == 0{
			fmt.Println("Fizz")
		} else if n%3 == 0{
			fmt.Println("Buzz")
		} else{
			fmt.Println(n)
		}

	}
}