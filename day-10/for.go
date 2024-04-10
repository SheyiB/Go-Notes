package main

import "fmt"

func main(){
	e := 1
	for e <=24{
		fmt.Println(e)
		e = e +1
	}

	for o := 0; o <10; o++ {
		fmt.Println(o)
	}

	for v := range 3{
		fmt.Println("range", v)
	}

	for { 
		fmt.Println("Printing till it crashes")
		break
	}

	for s := range 12{
		if s%2 == 0{
			continue
		}
		fmt.Println(s)
	}
}