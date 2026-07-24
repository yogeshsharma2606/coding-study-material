package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ------------------
// Helpers
// ------------------

func isVowel(c byte) bool {
	return strings.ContainsRune("aeiou", rune(c))
}

func isAllUpper(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) && !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

func isTitleCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	r := []rune(s)
	return unicode.IsUpper(r[0])
}

// ------------------
// Pig Latin Word
// ------------------

func pigLatinWord(word string) string {

	original := word
	word = strings.ToLower(word)

	// convert
	if isVowel(word[0]) {
		word = word + "yay"
	} else {
		for i := 0; i < len(word); i++ {
			if isVowel(word[i]) {
				word = word[i:] + word[:i] + "ay"
				break
			}
		}
	}

	// reapply casing
	if isAllUpper(original) {
		return strings.ToUpper(word)
	}

	if isTitleCase(original) {
		return strings.ToUpper(word[:1]) + word[1:]
	}

	return word
}

// ------------------
// Sentence
// ------------------

func pigLatinSentence(sentence string) string {
	words := strings.Fields(sentence)
	res := []string{}

	for _, w := range words {
		res = append(res, pigLatinWord(w))
	}

	return strings.Join(res, " ")
}

// ------------------
// Main
// ------------------

func main() {

	fmt.Println(pigLatinSentence("Pleased to meet you"))
	fmt.Println(pigLatinSentence("Do you speak Pig Latin"))
	fmt.Println(pigLatinSentence("Time flies when you are having fun"))
	fmt.Println(pigLatinSentence("A tree whose elements have at most two children is called a BINARY TREE"))
}
