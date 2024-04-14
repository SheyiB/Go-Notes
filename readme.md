# 🌟 Learning Golang Repository 🌟

This repository was created to **track my progress** as I embark on the exciting journey of learning the **Go programming language**. 🚀

I intentionally made the decision to start learning Golang on **April 1, 2024**. 🎉

## Decision to Learn Golang

## 📈 Progress

### Completed

### 🚧 In Progress

## 📝 Daily Learnings

This section will document my learnings day by day as I progress through the Go programming language.

### Day 1 (April 1, 2024)

- Learnt about the Go language. It was created by Google. Why? It was created to provide a programming language that is efficient, scalable and suitable
- How: it supports concurrency and parallelism improving performance.
- It makes use of light weight threads called Goroutines.
- It also includes garbage collection and allows low-level programming with ease
- It's simple to learn and read. It's also statically typed ensuring code safety
- It's also multiplatform and is compatible with a lot of cpu architectures
- It also features the Goformatter, and a testing framework
- Finally it's open source

- I had issues setting up Go on my Kali Linux. I'll try again tomorrow ...update, I go it! Thanks to Maxat Akbanov on Hashnode

- So with that being done, I printed my first `Jesus is Lord`, more like your typical "Hello world". I used fmt. I imported the `fmt` package and called it in what looked like a function and there was this package main at the top. Println is what was used to print the message. This language seems to combine a bit of java and c++ syntax, maybe just java sha. I ran it using `go run .`

- Forgot to add, before creating a `hello.go` file i ran `go mod init example/hello` to set up the module file which is like the package.json file, if you're coming from javascript

- I'm so gassed 🕺️🕺️🕺️🕺️🥳️🥳️🥳️!!! Goodnight guys

### Day 2 (April 2, 2024)

- I did not do much today, was busy working on a project, but still made little time
- Learnt about importing packages, checked a rsc.io for packages
- Worked with quote. Had issues with compiling, found a solution by setting my current user as root

### Day 3 (April 3, 2024)

- Today I revisited what I learnt yesterday and redid the example. Basically I imported a quote module and used one of it's features `.go()`. That returned a message which I printed.
- Next was to actually install the package so I ran `go mod tidy`. This installed the package and enabled my program to run.  These packages are seen in the `go.mod` and a more extensible one is seen in the `go.sum`. I think with how this relates to js it's like the `package.json` and `package.lock.json` files
- Furthermore I learnt how to create my own package which can be reused.
- Started by writing a function which took in a string and returned a string. Learnt about how to add types to function inputs and specify function return type
- Also learnt about how to print a formatted message(message that has a variable in it)
- Also learnt that *a function whose name starts with capital letter can be called by a function not in the same package* - This is the `exported name` concept
- In Go you can create a variable and use it at the same time like this `:=`

### Day 4 (April 4, 2024)

- Today I revisited what I learnt yesterday which was how to import a package you created. I had issues running it yesterday, but it worked today
- Today I learnt about error handling, how to check for an error case using an if statement and the create an error using the error module. Intersting that to create a new error in go, you say `errors.New()` unlike Js that says `new Error()`
- I also learnt how to check for errors and log them using the log module. It's intersting that in go when logging errors you can add a prefix and disable extra info from showing

### Day 5 (April 5, 2024)

- Today was a continuation of yesterday's lesson. I worked with the random function
- Learnt about Go slice, which is an array of dynamic size. So I created a slice of messages, stored in a variable and using the `len`(length) of this slice, generated a random integer within this value and returned the slice[index of randomly gen number]
- This randomly gen message then gets printed. Interestingly, Go uses the same method as python for getting the length of an array. Similar syntaxes everywhere, even for getting value at index.
- Ran into an error, in `go-slice` every item must end with a `comma(,)`. Interesting
- Although I've learnt this before i was reminded that in go: `functions with small letters cannot be exported`

### Day 6 (April 6, 2024)

- Today basically I created another function that took in a slice of names, looped through each name and returned greetings (the function created yesterday) for each name. The result was returned in a map
- I learnt about backward compatibility i.e rather than modifying your function, create a new one that builds on the existing function
- Also learn about the range tool, it returns an integer(postion of the item) and a value. Here I also learnt about the `Go blank` identifier which is used to ignore a value.
- Learnt about map. Which is more like the dictionary in Python. It's simply a key, value pair, and you can define the type like this `map[string]string` which simply means the key is a string and the value is a string
- I also learnt about a for loop, updating a map and creating a slice of values.

### Day 7 (April 7, 2024)

- Today I learnt about Writing tests in go
- Testing in go can be done using the go `testing module`
- In the test I did, I tested one of the perviously created function to ensure it returned a value.
- Test files end with`_test.go` and test functions always start with a Capital letter
- Test functions take in the pointer to the test package like this `(t *testing.T)`
- In writing the test, the function to be tested is called and the output is checked to ensure it is what is required. One variable used was the `want`.
- Tests that fail are logged using the `t.Fatalf()`
- I also leant about `Logical OR operators` used like this `||`
- Tests are ran using the `go test` command. You can also increase the details of the message shown using the `-v flag`

### Day 8 (April 8, 2024)

- Today I learnt about compiling modules
- You can compile using `go build` it creates an executable you can run independently
- Interstingly you can install this executable using `go install` (after exporting the path where you code get built, you can check this using `go list -f '{{.Target}}'` and export using `export PATH=$PATH:/path/to/your/install/directory`)
- Once this is done run `go install` with this, your compiled application can run anywhere on your system. You can open any directory and run your app. `Super COOOL!!!`

### Day 9 (April 9, 2024)

- Today I learnt about values types: string, integers, boolean, logical OR, AND.
- I also learnt about variables and constants and how to add types to them. It's the same a js
- I used the math module-sin.
-In Go, a variable can be empty but a constant cant

### Day 10 (April 10, 2024)

- Today I learnt for loops properly, apparently there are no while loops. Go employs the java/c++ style of for loops, using a start value, end and increment `i=1; i<10; i++`
- Also in go for loops, there are diffrent approaches, you can loop over a number, range, and even use the for loop without any condition to keep looping repeatedly. Here i also learnt about `break` and `continue`
- Furthermore, I learnt about if_else in go. You do not need to put your conditions in parenthesis, just write them after your if statment, and your put the corresponsding action in the `{}`
- Go uses the traditional if, else if and else
- I used these lessons to create a fizzbuzz program. I ran into an error. Apparently, when writing `else if`, the else if must be on the same line as the last `}` of the if statement or else if statement. With this i guess it's safe to say that there's some sort of pagiantion in go

### Day 11 (April 11, 2024)

- Todayyyyyyyyyyy, I learnt about building RESTful APIsssss!!! I'm so gassed
- I used the Go gin package
- I'll write more tomorrow, I'm behind schedule for today! Proud of my learnings so far
- So about yesterday I learnt about creating structs which are like interfaces but much more
- I also created endpoints for get, by id and post

### Day 12 (April 12, 2024)

- I learnt about arrays and slices, they're quite similar the major difference is array has fixed size while slice can be flexible
- You can check len of an array, cap, copy array.
- You can also create a slice using the `make` statement
- I learnt about switch, the syntax is quite similar to js. One additional in go is you might not put a condition value you can use the switch and in the case put your conditions, this would make the switch work like a conventional if else statement

### Day 13 (April 13, 2024)

- I learnt about maps which are like dictionaries in python. You can create maps by using the `make` keyword or `map`. I also learnt basic CRUD with maps
- I also learnt abour range, using range with a array returns the index and value, range with map, returns key and value, range with string returns the string position(index) and the unicode of the string

### Day 14 (April 14, 2024)

- I learnt about functions, adding return types, argument types.
- I also learnt how to return multiple values, and also accept arguments of variable length

## 🧠 Key Concepts & New Things Learned 🌱

Here are some nuggets I've picked up:

- **Short variable declaration**
- **Type inference**: function & variables
- **Print Statements**
- **Functions**
- **Error Handling and logging**
- **If statements**
- **Random Module**
- **Len (Get length of an array)**
- **Go Slice**
- **For loop**
- **Go Blank**
- **Backward Compatibility**
- **Go Maps**
- **Range**
- **Tests**
- **Logical Operators**
- **Compilation and Installation**
- **Data Types**
- **Variables and Constants**
- **Switch**

## 📚 Resources

### Books

- **None for now**

### 🌐 Websites

- [Golang Official Website](https://golang.org/)
- [A Tour of Go](https://tour.golang.org/)

**WAGMI! 🚀**

---
