package main

import (
	"fmt"
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

func randWord(arr []string) string {
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

func randChance(n float32) bool {
	if rand.Float32()*100 < n {
		return true
	}
	return false
}

func capitalize(s string) string {
	return string(string(unicode.ToUpper(rune(s[0]))) + s[1:])
}

func joinNounPhrases(phrases []nounPhrase, sep string) string {
	ret := ""
	for i, ph := range phrases {
		if i == 0 {
			ret += ph.String()
		} else {
			ret += fmt.Sprintf("%s%s", sep, ph)
		}
	}
	return ret
}
