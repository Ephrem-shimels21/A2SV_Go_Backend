package main

import (
	"fmt"
	"unicode"
)

func frequencyCounter(word string) map[string]int {
	var frequencyMap = make(map[string]int)
	for _, letter := range word {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			if unicode.IsLetter(letter) {
				letter = unicode.ToLower(letter)
			}
			frequencyMap[string(letter)]++

		}
	}
	return frequencyMap
}

func palindromeChecker(word string) bool {
	var reversedWord string
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	return reversedWord == word
}

func main() {
	fmt.Println("Enter a word to check its palindrome status: ")
	var inputWord string
	fmt.Scan(&inputWord)
	fmt.Printf("The word %s is a palindrome: %t\n", inputWord, palindromeChecker(inputWord))

	fmt.Println("Enter a word to check its frequency: ")
	var word string
	fmt.Scan(&word)
	fmt.Printf("The frequency of the word %s is: %v\n", word, frequencyCounter(word))
}
