# Go

A repo containing all about the GO during my learnings

Hope this makes go language easy to learn.  😃😃


###  [GO OFFICIAL WEBSITE](https://golang.org/)

### My environment Setup

1. VSCODE
2. GO official vscode extension
3. A windows 10 OS😁😁


# Basics

Lets start with the most basic hello world program

`hello.go`

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello world")
}
```

Make sure to use double quotes and not single quotes

Now some of the questions that may come in our mind after seeing the program are:-

1. **How do we run the project?**
   
   - Go to the program directory and run `go run hello.go`

    - Make sure to install the go compiler from the official website mentioned above.

    ### GO CLI Commands

    1. `go build` - compiles a bunch of source code files (it does not execute it but just create a binary eg. for windows it will be an exe file)
    2. `go run` - compiles or ececutes one or more files
    3. `go fmt` - formats all the code in each file in the current directory
    4. `go install` - compiles and install a package
    5. `go get` - download the raw code of someone else package
    6. `go test` - run any test associated with the current project

<br/>

2. **What does `package main` means?**

    - Package is like workspace of common source code file for a project 

    - package == project ==workspace

    - every file of same package must contain same name package

    - Now there are two types of package `executable` and `reusable`
    - `Executable` generates a file that we can run
    - `reusable` is used as "helpers".It is good place to put reusable logic
    - creating final binary using build works only for `executable` type package
    - So to know if we are making executable package we need to specify `main` as package name
    - If we use any other package name other than main and then run ` go build filename` then it wont spit out the executable file or the binary
    - Everytime we make an executable package it should have a **functioon**(func) called **`main`**


3. **What does `import "fmt"` means?**

    - The import statement is used to get all code inside some other package
    - here `fmt` is used for format library usually for console output and debug purpose
    - there are other packages like **crypto, io , debug , math ,ecoding** etc
    - `fmt` is a standard library package
    - we can find the list of all standard package in this link [standard packages link](https://pkg.go.dev/std)

4. **What is the `func()`?**

    - func refers to functions similar to other programming language
    - `func main(){}`
    - `func` tells we are about to declare a function
    - `main` sets the name of the function
    - `()` - here we can pass the list of arguments to the functioon
    - `{}` -> function body

5. **How is the `hello.go` file organized?**

    - `package main` -> package declaration
    - `import "fmt"` -> import other packages that we need
    - `func main(){...}` -> declare fnctions 
    - This is the same pattern used in every go file 


# Our First Project - Deck Of Playing Cards

Here in this program we simulate the deck of playing cards
Using this we will be learning the basics of this language

**Features**:- 

1. newDeck -> create a list of playing cards essentially an array of strings
2. print -> log out the contents of a deck of cards
3. shuffle -> shuffle all the cards in a deck
4. deal -> create a hand of cards
5. saveToFile -> save list of cards to a file in local machine
6. newDeckFromFile -> load a list of cards from the local machine 

Create a new project directory

## Declaring a variable 

Go is a static typed language

```go
package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	fmt.Println(card)
}
```

- `var` tells that its variable 
- `card` tells the name of the variable
- `string` tells that only string will be stored in this variable

    ## Basic Data Types in GO

    1. bool
    2. string
    3. int
    4. float64

we can also use `card:= "Ace of Cards"` here the compiler get to know about the type of variable being stored

This syntax is used only while first initializing

```go
package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of spades"
	fmt.Println(card)
}
```

Note :-  Variables can be initialized outside of a function, but cannot be assigned a variable.