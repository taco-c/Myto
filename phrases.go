package main

import (
	"fmt"
	"strings"
)

type noun struct {
	a  string
	sg string
	pl string
}

type adjectivePhrase struct {
	adjective string
	adverb    string
}

type nounPhrase struct {
	qualifier string
	compound  string
	noun      string
	of        string
}

func (ph nounPhrase) String() string {
	str := strings.Title(ph.noun)
	if ph.compound != "" {
		str = fmt.Sprintf("%s-%s",
			strings.Title(ph.compound),
			str,
		)
	}
	if ph.qualifier != "" {
		str = fmt.Sprintf("%s %s",
			ph.qualifier,
			str,
		)
	}
	if ph.of != "" {
		str = fmt.Sprintf("%s of %s",
			str,
			strings.Title(ph.of),
		)
	}
	return str
}

func randActorPhrase() nounPhrase {
	noun := randWord(actors)
	return randDescriptorsToNoun(noun)
}

func randDescriptorsToNoun(noun string) nounPhrase {
	var qualifier, of, compound string
	var chance float32 = 40.0

	// forms := choiceList(actoractors)
	// noun := noun{
	// 	a:  forms[0],
	// 	sg: forms[1],
	// 	pl: forms[2],
	// }
	if randChance(chance) {
		qualifier = randWord(adjectives)
		chance /= 3
	}
	if randChance(chance) {
		of = randWord(abstractStuff)
	}
	if randChance(5) {
		compound = randWord(items)
	}

	return nounPhrase{qualifier, compound, noun, of}
}

func randLocationPhrase() nounPhrase {
	var chance float32 = 50.0
	var ph = nounPhrase{
		noun: randWord(locations),
	}
	if randChance(chance) {
		ph.of = randWord(abstractStuff)
		chance /= 2
	}
	if randChance(chance) {
		ph.qualifier = randWord(adjectives)
	}

	return ph
}

func randItemPhrase() nounPhrase {
	var chance float32 = 75.0
	var ph = nounPhrase{
		noun: randWord(items),
	}
	if randChance(chance) {
		ph.of = randWord(abstractStuff)
		chance /= 4
	}
	if randChance(chance) {
		ph.qualifier = randWord(adjectives)
	}

	return ph
}

type verb struct {
	inf  string
	sg3  string
	past string
	ing  string
}

type verbPhrase struct {
	verb    string
	nom     nounPhrase
	acc     nounPhrase
	dat     nounPhrase
	instr   nounPhrase
	loc     nounPhrase
	locPrep string
	adverb  string
}

func (ph verbPhrase) String() string {
	var adverb string
	var str string = ph.verb
	// if ph.adverb == "" || randBoolChance(25) {
	// 	adverb = choiceString(adverbs)
	// } else {
	// 	adverb = ph.adverb
	// }
	adverb = ph.adverb
	if adverb != "" {
		if ph.acc != (nounPhrase{}) {
			str = fmt.Sprintf("%s %s", adverb, str)
		} else {
			str = fmt.Sprintf("%s %s", str, adverb)
		}
	}
	if ph.nom != (nounPhrase{}) {
		str = fmt.Sprintf("the %s %s", ph.nom, str)
	}
	if ph.acc != (nounPhrase{}) {
		str = fmt.Sprintf("%s the %s", str, ph.acc)
	}
	if ph.dat != (nounPhrase{}) {
		str = fmt.Sprintf("%s to the %s", str, ph.dat)
	}
	if ph.instr != (nounPhrase{}) {
		if randChance(75) {
			str = fmt.Sprintf("%s with the %s", str, ph.instr)
		} else {
			str = fmt.Sprintf("with the %s, %s", ph.instr, str)
		}
	}
	if ph.loc != (nounPhrase{}) {
		if randChance(75) {
			str = fmt.Sprintf("%s %s the %s", str, ph.locPrep, ph.loc)
		} else {
			str = fmt.Sprintf("%s the %s, %s", ph.locPrep, ph.loc, str)
		}
	}
	return capitalize(str)
}

func randActorVerbPhrase(nom nounPhrase) verbPhrase {
	var verb, locPrep, adverb string
	var acc, dat, instr, loc nounPhrase
	var chance float32 = 50.0

	if randChance(50) {
		verb = randWord(transVerbs)
		for ok := true; ok; ok = (acc == nom) {
			acc = randActorPhrase()
		}
	} else {
		verb = randWord(intransVerbs)
	}
	// if randBoolChance(chance) {
	// 	for ok := true; ok; ok = dat == nom || dat == acc {
	// 		dat = randNounPhrase()
	// 	}
	// }
	if randChance(chance) {
		if randChance(25) {
			loc = randActorPhrase()
			locPrep = "by"
		} else {
			loc = randLocationPhrase()
			locPrep = "at"
		}
		chance /= 2
	}
	if randChance(chance) {
		if randChance(50) {
			for ok := true; ok; ok = instr == nom || instr == acc || instr == dat {
				instr = randActorPhrase()
			}
		} else {
			instr = randItemPhrase()
		}
	}
	if randChance(25) {
		adverb = randWord(adverbs)
	}
	return verbPhrase{verb, nom, acc, dat, instr, loc, locPrep, adverb}
}

func randVerbPhraseFromContent(localActors, localItems, localLocations []nounPhrase) verbPhrase {
	var verb, locPrep, adverb string
	var nom, acc, dat, instr, loc nounPhrase
	var chance float32 = 50.0

	nom = choice(localActors)
	if randChance(50) && len(localActors) > 1 {
		verb = randWord(transVerbs)
		for ok := true; ok; ok = acc == nom {
			acc = choice(localActors)
		}
	} else {
		verb = randWord(intransVerbs)
	}
	// if randBoolChance(chance) {
	// 	for ok := true; ok; ok = dat == nom || dat == acc {
	// 		dat = randNounPhrase()
	// 	}
	// }
	if randChance(chance) && len(localLocations) > 0 {
		if randChance(25) && len(localActors) > 2 {
			loc = choice(localActors)
			locPrep = "by"
		} else {
			loc = choice(localLocations)
			locPrep = "at"
		}
		chance /= 2
	}
	if randChance(chance / 2) {
		if randChance(25) && len(localActors) > 2 {
			for ok := true; ok; ok = instr == nom || instr == acc || instr == dat {
				instr = choice(localActors)
			}
		} else if len(localItems) > 0 {
			instr = choice(localItems)
		} else {
			instr = nounPhrase{}
		}
	}
	if randChance(0) {
		adverb = randWord(adverbs)
	}
	return verbPhrase{verb, nom, acc, dat, instr, loc, locPrep, adverb}
}

func randVerbPhrase() verbPhrase {
	var actor nounPhrase = randActorPhrase()
	return randActorVerbPhrase(actor)
}
