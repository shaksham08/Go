# Go

A repo containing all about the GO during my learnings

Hope this makes go language easy to learn. ðð

### [GO OFFICIAL WEBSITE](https://golang.org/)

### My environment Setup

1. VSCODE
2. GO official vscode extension
3. A win 11 OSðð

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

Make sure to use double quotes and not single quotes.

Single quotes are used for single characters, called **runes**

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
   - Everytime we make an executable package it should have a **function**(func) called **`main`**

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

This syntax with `:` is used only while first initializing

```go
package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of spades"
	fmt.Println(card)
}
```

Note :- We can initialize a variable outside of a function, we just can't assign a value to it.

```go
package main

import "fmt"

var deckSize int

func main() {
  deckSize = 50
  fmt.Println(deckSize)
}
```

```go
package main

import "fmt"

var deckSize int
deckSize = 101

func main() {
  deckSize = 50
  fmt.Println(deckSize)
}
```

This above program wont work

## Functions and return Types

```go
package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

```

Here `newCard() string` tells that it will return string type of data

This return type can be anything with whatever we are returning

Note:- **Files in the same package can freely call functions defined in other files.**

- main.go

```go
package main

func main() {
    printState()
}
```

- state.go

```go
package main

import "fmt"

func printState() {
    fmt.Println("California")
}
```

Running both the files using `go run main.go state.go` will perfectly run

if we run only main.go then it would give error as `undefined: printState`

## Array and Slice

Now this `newCard` function will return more than one card so we need array

Array :- Fixed length of things
Slice :- An Array that can grow or shrink

These both should be defined with some data type so that we can store same type of data in this.

slice

```go
package main

import "fmt"

func main() {

	//slice
	card := []string{"Ace of diamonds", newCard()}
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
```

Now we want to insert something in the slice so we use append

Every element in slice must be of same type

```go

	cards = append(cards, "Six of Spades")
```

Here append just returns new slice which

Now to iterate and print the cards we use

```go
//iterating over the slice to print elements of it
	for i, card := range cards {
		fmt.Println(i, card)
	}
```

```go
package main

import "fmt"

func main() {

	//slice
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterating over the slice to print elements of it
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
```

here for iterating we are using for

index = index of this element in the array

card = current card

range cards = Take the slice of cards and loop over it

we are using `:=` because everytime we are iterating we are getting new variables of i and card and we know this syntax is only used when we first initialize the variable in go

Note:- This will not work below code

```go
for index, card := range cards {
		fmt.Println(card)
	}
```

here we will get error as `index declared but not used`

## Object oriented approach in GO

Note:- Go is not an OO language
Here there is no idea of classes

![OOApproach](./images/ooapproach.png)

![OOAPROACH2](./images/ooapproach2.png)

![OOAPROACH3](./images/ooapproach3.png)

![folderstructure](./images/folderstucture.png)

Now we want to create a deck type

deck.go

```go
package main

// create a new type of deck
//which is a slice of string
type deck []string

//this says it is equal to strings
```

main.go

```go
package main

import "fmt"

func main() {

	//slice
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterating over the slice to print elements of it
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

```

Now we add our own reciever function

```go
package main

import "fmt"

// create a new type of deck
//which is a slice of string
type deck []string

//this says it is equal to strings

//here inside func we are putting reciver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

```

main.go

```go
package main

import "fmt"

func main() {

	//slice
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
```

`func (d deck) print()` - this is a reciever functions which tells any veriable of type deck can access to the print method

Here d is the actual copy of the deck we are working with is available in the function as a variable called 'd'

deck -> every variable of type deck can call this function on itself

![recieverfunction](images/recieverfunction.png)

Note:- In Go we dont use `this` of `self`

Now we add a new function to create a new Deck

```go
// we dont need revciver here as still we dont have any deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}
```

Here \_ means we dont need to use these variables anywhere

## Slice Range Index

_fruits[startIndexIncluding : upToNotIncluding]_

_fruits[0:2]_

## Return multiple values from functions

```go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

## GO testing

Go testing is not mocha,jasmine or selenium etc

to make a test,create a file ending with \_test.go

`deck_test.go`

to run all tests in a package , run the command

`go test`

## Struct

Data structure , collection of properties that are related together

```go
package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
}
```

we can also use different syntax

```go
package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
}
```

Another way of defining stucts is

```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	alex.firstName = "alex"
	alex.lastName = "anderson"
}

```

one more example

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	fmt.Printf("%+v", jim)
}
```

### Reciever functions with struct

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```

Now if we try to update it we wont be able to do

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.updateName("shaksham")
	jim.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```

## Pass by Value

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("shaksham")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```

Here

`&variable` -> give me the memory address of the value this variable is pointing at

`*pointer` -> give me the value this memory address is pointing at

## Pointer Shortcut

```go
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.updateName("shaksham")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```

```go

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.updateName("shaksham")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

```

## Map in Go

In go a map is just a key value pair

- All key or value should be of same type

```go
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	fmt.Println(colors)
}
```

Another way

```go
package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	var colors map[string]string

	fmt.Println(colors)
}
```

```go
package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors,"white")

	fmt.Println(colors)
}

```

Iterating over Maps

```go
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
```

## Interfaces

We know that every value has a type and every function has d every function has to be specify the type of its argument

so it means that every function we ever write has to be written to accommodate different types even if the logic in it it identical??

```go
package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Print(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi There!"
}

func (eb spanishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi There!"
}
```

Interfaces are not generic types

Interfaces are implicit

Interfaces are a contract to help us manage types

interfaces are tough

## The HTTP Package

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

```
