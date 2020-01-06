package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func printRandomStory(storyLength, verseLength int) {
	// var newActor, actor, previousActor nounPhrase
	// var parents [2]nounPhrase
	var newActorChance float32 = 50.0
	var newItemChance float32 = 20.0
	var newLocationChance float32 = 20.0
	var localActors = []nounPhrase{}
	for i := 0; i < 1+rand.Intn(2); i++ {
		localActors = append(localActors, randActorPhrase())
	}
	var localItems = []nounPhrase{}
	var localLocations = []nounPhrase{}

	// if len(localActors) == 2 {
	// 	fmt.Printf("In the beginning there is the %s and the %s.\n",
	// 		localActors[0],
	// 		localActors[1],
	// 	)
	// } else {
	// 	fmt.Printf("In the beginning there is the %s.\n",
	// 		localActors[0],
	// 	)
	// }

	for i := 0; i < storyLength; i++ {
		for j := 0; j < verseLength; j++ {
			// if i == 0 && j == 0 {
			// 	continue // Compensation for the opening line above.
			// }
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
			// previousActor = actor
			// if randChance(newActorChance) {
			// 	newActorChance -= newActorChance / 4
			// 	newActor = randActorPhrase()
			// 	// Parents
			// 	if len(localActors) <= 1 || randChance(5) {
			// 		fmt.Printf("Then the %s is born from the %s.\n",
			// 			newActor,
			// 			choice(localActors),
			// 		)
			// 	} else {
			// 		parents[0] = choice(localActors)
			// 		for ok := true; ok; ok = parents[0] == parents[1] {
			// 			parents[1] = choice(localActors)
			// 		}
			// 		fmt.Printf("Then the %s is born from the %s, and the %s.\n",
			// 			newActor,
			// 			parents[0],
			// 			parents[1],
			// 		)
			// 	}
			// 	j++
			// 	localActors = append(localActors, newActor)
			// 	actor = newActor
			// } else {
			// 	if randChance(25) && previousActor != (nounPhrase{}) {
			// 		actor = previousActor
			// 	} else {
			// 		actor = choice(localActors)
			// 	}
			// }

			// TODO: And then xxx (location) is discovered/built.

			fmt.Printf("%s.\n", randVerbPhraseFromContent(localActors, localItems, localLocations))

			// if n == 4 {
			// 	fmt.Printf("The %s thinks of %s %s the %s.\n",
			// 		actor,
			// 		choiceString(abstractStuff),
			// 		choiceString(prepositions),
			// 		choice(items),
			// 	)
			// } else if n == 5 {
			// 	fmt.Printf("Then the %s %s, making the %s %s.\n",
			// 		actor,
			// 		choiceString(intransVerbs),
			// 		choice(items),
			// 		choiceString(adjectives),
			// 	)
			// }
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

	// for _, item := range items {
	// 	fmt.Printf("The %s is the symbol of %s.\n",
	// 		item.string(),
	// 		choiceString(abstractStuff),
	// 	)
	// }
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
