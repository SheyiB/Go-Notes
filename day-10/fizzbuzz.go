package main

import "fmt"

func main(){
	for n:= range 100{
		if n%5 == 0 && n %3 == 0{
			fmt.Println("Fizzbuzz")
		}
		else if n%5 == o{
			fmt.Println("Fizz")
		}
		else if n%3 == 0{
			fmt.Println("Buzz")
		}
		else{
			fmt.Println(n)
		}

	}
}