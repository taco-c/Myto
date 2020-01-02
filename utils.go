package main

import (
	"math/rand"
	"unicode"
)

func choice(arr []nounPhrase) nounPhrase {
	l := len(arr)
	if l == 0 {
		return nounPhrase{}
	}
	return arr[rand.Intn(l)]
}

func choiceString(arr []string) string {
	l := len(arr)
	if l == 0 {
		return ""
	}
	return arr[rand.Intn(l)]
}

func choiceList(arr [][]string) []string {
	l := len(arr)
	if l == 0 {
		return make([]string, 0)
	}
	return arr[rand.Intn(l)]
}

func randBoolChance(n int) bool {
	if rand.Intn(100) < n {
		return true
	}
	return false
}

func capitalize(s string) string {
	return string(string(unicode.ToUpper(rune(s[0]))) + s[1:])
}
