package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "First input teste iii"
	wordToSearch := "i"
	frequencyOfWord := strings.Count(input, wordToSearch)
	fmt.Printf("The word %s have appeared %v times \n", wordToSearch, frequencyOfWord)
}
