package main

import "fmt"

func main() {
	first := "First: line 1\nline 2"
	second := `Second: line 1\nline 2`
	third := "Многоязычный string"

	fmt.Println("first      = ", first)
	fmt.Println("second     = ", second)
	fmt.Println("len(third) = ", len(third))

	stringOperations()
}

func stringOperations() {
	first := "firststring"
	second := "secondstring"
	space := " "
	word := "приветмир"
	runeWord := []rune(word)

	fmt.Println("\n--------------------------------------------")

	fmt.Println("first + space + second  = ", first + space + second)
	fmt.Println("first[1]                = ", first[1])
	fmt.Println("string(first[1])        = ", string(first[1]))
	fmt.Println("first[3:]               = ", first[3:])
	fmt.Println("first[:7]               = ", first[:7])
	fmt.Println("first[2:7]              = ", first[2:7])

	fmt.Println("word[1]            = ", word[1])
	fmt.Println("word[3:]           = ", word[3:])
	fmt.Println("word[:7]           = ", word[:7])
	fmt.Println("word[2:7]          = ", word[2:7])

	fmt.Println("word[1]            = ", runeWord[1])
	fmt.Println("word[3:]           = ", runeWord[3:])
	fmt.Println("word[:7]           = ", runeWord[:7])
	fmt.Println("word[2:7]          = ", runeWord[2:7])

	fmt.Println("\nstring(word[1])  = ", string(runeWord[1]))
	fmt.Println("string(word[3:])   = ", string(runeWord[3:]))
	fmt.Println("string(word[:7])   = ", string(runeWord[:7]))
	fmt.Println("string(word[2:7])  = ", string(runeWord[2:7]))
}