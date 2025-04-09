package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// remove punctuation from string
func removePunctuation(s string) string {
	var sb strings.Builder
	for  _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			sb.WriteRune(unicode.ToLower(c))
		}
	}
	return sb.String()
}

func is_palindrome(word string) bool{
	cleanedWord := removePunctuation(word)

	for i := 0; i <= (len(cleanedWord) / 2); i+=1{
		if cleanedWord[len(cleanedWord) - i - 1] != cleanedWord[i] {
			return false
		}
	}
	return true
}

func word_frequency(phrase string) map[string]int {
	dictionary := map[string]int{}
	words := strings.Split(strings.TrimSpace(phrase), " ")
	
	for _, word := range words {
		word = removePunctuation(strings.ToLower(word))
        if dictionary[word] != 0 {
            dictionary[word]++
        } else {
            dictionary[word] = 1
        }
	}
	return dictionary
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the word: ")
	phrase, _ := reader.ReadString('\n')
	
	fmt.Println("1. check for palindrome")
	fmt.Println("2. Count word frequency")
	fmt.Print("Which funcation do you select: ")

	selection, _ := reader.ReadString('\n')
	if strings.TrimSpace(selection) == "1" {
		fmt.Println(is_palindrome(phrase))
	}else{
		fmt.Println(word_frequency(phrase))
	}
	
}