package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func printRandomStory() {
	var localActors = []nounPhrase{
		randNounPhrase(),
		randNounPhrase(),
	}

	fmt.Printf("In the beginning there is the %s and the %s.\n",
		localActors[0],
		localActors[1],
	)

	var localItems = []nounPhrase{
		randItemPhrase(),
		// randItemPhrase(),
		// randItemPhrase(),
	}

	var localLocations = []nounPhrase{
		randLocationPhrase(),
		// randLocationPhrase(),
		// randLocationPhrase(),
	}

	var newActor, actor, previousActor nounPhrase
	var newActorChance int = 25
	var storySize int = 1 + rand.Intn(1)
	var verseSize int = 4 + rand.Intn(16)
	var n int
	var options []verbPhrase
	var parents [2]nounPhrase
	for i := 0; i < storySize; i++ {
		for j := 0; j < verseSize; j++ {
			if i == 0 && j == 0 {
				j++
				continue
			}
			previousActor = actor
			if randBoolChance(newActorChance) {
				newActorChance -= newActorChance / 4
				if newActorChance <= 0 {
					newActorChance = 0
				}
				newActor = randNounPhrase()
				// Parents
				if len(localActors) >= 2 || randBoolChance(95) {
					parents[0] = choice(localActors)
					for ok := true; ok; ok = parents[0] == parents[1] {
						parents[1] = choice(localActors)
					}
					fmt.Printf("Then the %s is born from the %s, and the %s.\n",
						newActor,
						parents[0],
						parents[1],
					)

				} else {
					fmt.Printf("Then the %s is born from the %s.\n",
						newActor,
						choice(localActors),
					)
				}
				localActors = append(localActors, newActor)
				j++
				actor = newActor
			} else {
				if randBoolChance(25) && previousActor != (nounPhrase{}) {
					actor = previousActor
				} else {
					actor = choice(localActors)
				}
			}
			// TODO: And then xxx (location) is discovered/built.

			options = []verbPhrase{
				// Intransitive verbs
				verbPhrase{
					verb: choiceString(intransVerbs),
					nom:  actor,
				},
				verbPhrase{
					verb:  choiceString(intransVerbs),
					nom:   actor,
					instr: choice(localActors),
				},
				verbPhrase{
					verb:  choiceString(intransVerbs),
					nom:   actor,
					instr: choice(localItems),
				},
				verbPhrase{
					verb:    choiceString(intransVerbs),
					nom:     actor,
					loc:     choice(localLocations),
					locPrep: "at",
				},
				// Transitive verbs
				verbPhrase{
					verb: choiceString(transVerbs),
					nom:  actor,
					acc:  choice(localActors),
				},
				verbPhrase{
					verb:  choiceString(transVerbs),
					nom:   actor,
					acc:   choice(localActors),
					instr: choice(localActors),
				},
				verbPhrase{
					verb:    choiceString(transVerbs),
					nom:     actor,
					acc:     choice(localActors),
					instr:   choice(localActors),
					loc:     nounPhrase{noun: "hand", qualifier: "its"},
					locPrep: "in",
				},
				verbPhrase{
					verb:  choiceString(transVerbs),
					nom:   actor,
					acc:   choice(localActors),
					instr: choice(localItems),
				},
				verbPhrase{
					verb:    choiceString(transVerbs),
					nom:     actor,
					acc:     choice(localActors),
					loc:     choice(localLocations),
					locPrep: "at",
				},
				verbPhrase{
					verb:    choiceString(transVerbs),
					nom:     actor,
					acc:     choice(localActors),
					loc:     choice(localLocations),
					locPrep: "at",
					instr:   choice(localActors),
				},
				verbPhrase{
					verb:    choiceString(transVerbs),
					nom:     actor,
					acc:     choice(localActors),
					loc:     choice(localLocations),
					locPrep: "at",
					instr:   choice(localItems),
				},
				verbPhrase{
					verb: "gives",
					nom:  actor,
					acc:  choice(localItems),
					dat:  choice(localActors),
				},
				verbPhrase{
					verb: "feels",
					nom:  actor,
					acc:  nounPhrase{noun: choiceString(abstractStuff)},
				},
				verbPhrase{
					verb: "thinks of",
					nom:  actor,
					acc:  nounPhrase{noun: choiceString(abstractStuff)},
				},
			}

			var ph verbPhrase
			for ok := true; ok; ok = ph.nom == ph.acc || ph.nom == ph.dat || ph.nom == ph.instr || ph.nom == ph.loc {
				n = rand.Intn(len(options))
				ph = options[n]
			}
			// fmt.Printf("%s.\n", ph.string())
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
		fmt.Printf("%2d %s\n", i+1, capitalize(item.String()))
	}

	fmt.Println("\nItems:")
	for i, item := range localItems {
		fmt.Printf("%2d %s\n", i+1, capitalize(item.String()))
	}

	fmt.Println("\nLocations:")
	for i, item := range localLocations {
		fmt.Printf("%2d %s\n", i+1, capitalize(item.String()))
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
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	if *doPrintStory {
		printRandomStory()
	} else if *doPrintSentence {
		for i := 0; i < 1; i++ {
			ph := randVerbPhrase()
			fmt.Printf("%s.\n", ph)
		}
	} else {
		fmt.Printf("Usage of %v:\n", os.Args[0])
		flag.PrintDefaults()
	}

}
