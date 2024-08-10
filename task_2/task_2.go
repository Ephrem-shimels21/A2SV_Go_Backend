package main

import (
	"fmt"
	"unicode"
)


func frequencyCounter(text string) map[string]int {
	frequencyMap := make(map[string]int)
	var word string
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			word += string(unicode.ToLower(r))
		} else if word != "" {
			frequencyMap[word]++
			word = ""
		}
	}
	if word != "" {
		frequencyMap[word]++
	}
	return frequencyMap
}

func palindromeChecker(word string) bool {
	var filteredWord string
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			filteredWord += string(unicode.ToLower(r))
		}
	}
	var reversedWord string
	for i := len(filteredWord) - 1; i >= 0; i-- {
		reversedWord += string(filteredWord[i])
	}
	return reversedWord == filteredWord
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
