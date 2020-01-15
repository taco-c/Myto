package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func printRandomStory(storyLength, verseLength int) {
	var newActorChance float32 = 50.0
	var newItemChance float32 = 20.0
	var newLocationChance float32 = 20.0
	var localActors = []nounPhrase{}
	for i := 0; i < 1+rand.Intn(2); i++ {
		localActors = append(localActors, randActorPhrase())
	}
	var localItems = []nounPhrase{}
	var localLocations = []nounPhrase{}

	fmt.Printf("In the beginning there is the %s.\n",
		joinNounPhrases(localActors, " and the "),
	)

	for i := 0; i < storyLength; i++ {
		for j := 0; j < verseLength; j++ {
			// Compensation for the opening line above.
			if i == 0 && j == 0 {
				continue
			}
			if randChance(newActorChance) {
				localActors = append(localActors, randActorPhrase())
				newActorChance *= 0.75
			}
			if randChance(newItemChance) {
				localItems = append(localItems, randItemPhrase())
				newItemChance *= 0.75
			}
			if randChance(newLocationChance) {
				localLocations = append(localLocations, randLocationPhrase())
				newLocationChance *= 0.75
			}
			fmt.Printf("%s.\n", randVerbPhraseFromContent(localActors, localItems, localLocations))
		}
		fmt.Println()
	}
	fmt.Println("\nActors:")
	for i, item := range localActors {
		fmt.Printf("%2d. %s\n", i+1, capitalize(item.String()))
	}
	fmt.Println("\nItems:")
	for i, item := range localItems {
		fmt.Printf("%2d. %s\n", i+1, capitalize(item.String()))
	}
	fmt.Println("\nLocations:")
	for i, item := range localLocations {
		fmt.Printf("%2d. %s\n", i+1, capitalize(item.String()))
	}
}

func main() {
	// Deal with flags
	var doPrintStory *bool = flag.Bool("m", false, "Print a mythological story")
	var doPrintSentence *bool = flag.Bool("s", false, "Print a random sentence")
	var length *int = flag.Int("l", 1, "Length of mythology or number of random sentences to print")
	var verseLength *int = flag.Int("v", 8, "Length of the verses of the mythology")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *doPrintStory {
		printRandomStory(*length, *verseLength)
	} else if *doPrintSentence {
		for i := 0; i < *length; i++ {
			fmt.Printf("%s.\n", randVerbPhrase())
		}
	} else {
		fmt.Printf("Usage of %v:\n", os.Args[0])
		flag.PrintDefaults()
	}

}
