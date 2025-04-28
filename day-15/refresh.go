package main

import "fmt"

func main() {
	// Intro
	fmt.Println("Back at it!")

	// Basic Variables
	name := "Elijah"
	var surname string = "Banjo"
	var age int = 24

	fmt.Println("Welcome", name, surname)
	fmt.Println("You are", age, "?")
	age = 25
	fmt.Println("You are turning",age,"?")

	const Lord string = "Jesus"
	const spouse int = 1


	fmt.Println("Your master is", Lord,"!")
	// ERROR!
	// Lord = "mammon"
	fmt.Println("Therefore, you can only have", spouse, "spouse")
	//ERROR
	//spouse = 3
	//const spouse int = 3

	//For
	//Basic
	i := 5
	fmt.Println("Five to one countdown to your birthday...")
	for i > 0{
		fmt.Println(i)
		i = i -1
	}
	fmt.Println("Happy Birthday!")

	//Condition
	fmt.Println("Again...")
	for i := 5; i > 0; i--{
		fmt.Println(i)
	}
	fmt.Println("Happy Birthday!")
	//i is a variable hence the reason why we can reuse it

	//Range
	//Now let's give thanks
	for i := range 24{
		fmt.Println("We thank God for year", i+1)
	}
	fmt.Println("We thank God for what He will do in year 25!")

	//Condition-less
	for {
		fmt.Println("Loop!")
		//if you don't do this, oti lor
		break
	}

	//This will print all odd numbers
	fmt.Println("Odd Numbers")
	for n := range 6{
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//This will print all even numbers
	fmt.Println("Even Numbers")
	for n := range 20{
		if n%2 != 0{
			continue
		}
		fmt.Println(n)
	}


}