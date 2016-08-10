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
	first := "first"
	second := "second"
	space := " "
	word := "приветмир"
	runeWord := []rune(word)

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("first + space + second  = ", first + space + second)
	first += space + second;
	fmt.Println("first += space + second = ", first)
	fmt.Println("first[1]                = ", first[1])
	fmt.Println("string(first[1])        = ", string(first[1]))
	fmt.Println("first[3:]               = ", first[3:])
	fmt.Println("first[:7]               = ", first[:7])
	fmt.Println("first[2:7]              = ", first[2:7], "\n")

	fmt.Println("ukrainian[1]            = ", word[1])
	fmt.Println("ukrainian[3:]           = ", word[3:])
	fmt.Println("ukrainian[:7]           = ", word[:7])
	fmt.Println("ukrainian[2:7]          = ", word[2:7], "\n")

	fmt.Println("runeUkrainian[1]            = ", runeWord[1])
	fmt.Println("runeUkrainian[3:]           = ", runeWord[3:])
	fmt.Println("runeUkrainian[:7]           = ", runeWord[:7])
	fmt.Println("runeUkrainian[2:7]          = ", runeWord[2:7])

	fmt.Println("\nstring(runeUkrainian[1])  = ", string(runeWord[1]))
	fmt.Println("string(runeUkrainian[3:])   = ", string(runeWord[3:]))
	fmt.Println("string(runeUkrainian[:7])   = ", string(runeWord[:7]))
	fmt.Println("string(runeUkrainian[2:7])  = ", string(runeWord[2:7]), "\n")
}