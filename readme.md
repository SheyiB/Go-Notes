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
- Although I've learnt this before i was reminded that in go:
- `functions with small letters cannot be exported`

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

## 📚 Resources

### Books

- **None for now**

### 🌐 Websites

- [Golang Official Website](https://golang.org/)
- [A Tour of Go](https://tour.golang.org/)

**WAGMI! 🚀**

---
