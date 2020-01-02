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

func randNounPhrase() nounPhrase {
	noun := choiceString(actors)
	return randDescriptorsToNoun(noun)
}

func randDescriptorsToNoun(noun string) nounPhrase {
	var qualifier, of, compound string
	var chance int = 40

	// forms := choiceList(actoractors)
	// noun := noun{
	// 	a:  forms[0],
	// 	sg: forms[1],
	// 	pl: forms[2],
	// }
	if randBoolChance(chance) {
		qualifier = choiceString(adjectives)
		chance /= 3
	}
	if randBoolChance(chance) {
		of = choiceString(abstractStuff)
	}
	if randBoolChance(5) {
		compound = choiceString(items)
	}

	return nounPhrase{qualifier, compound, noun, of}
}

func randLocationPhrase() nounPhrase {
	var ph = nounPhrase{
		noun: choiceString(locations),
	}
	var chance int = 50
	if randBoolChance(chance) {
		ph.of = choiceString(abstractStuff)
		chance /= 2
	}
	if randBoolChance(chance) {
		ph.qualifier = choiceString(adjectives)
	}

	return ph
}

func randItemPhrase() nounPhrase {
	var ph = nounPhrase{
		noun: choiceString(items),
	}
	var chance int = 75
	if randBoolChance(chance) {
		ph.of = choiceString(abstractStuff)
		chance /= 4
	}
	if randBoolChance(chance) {
		ph.qualifier = choiceString(adjectives)
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
		if randBoolChance(75) {
			str = fmt.Sprintf("%s with the %s", str, ph.instr)
		} else {
			str = fmt.Sprintf("with the %s, %s", ph.instr, str)
		}
	}
	if ph.loc != (nounPhrase{}) {
		if randBoolChance(75) {
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
	var chance int = 50

	if randBoolChance(50) {
		verb = choiceString(transVerbs)
		for ok := true; ok; ok = (acc == nom) {
			acc = randNounPhrase()
		}
	} else {
		verb = choiceString(intransVerbs)
	}
	// if randBoolChance(chance) {
	// 	for ok := true; ok; ok = dat == nom || dat == acc {
	// 		dat = randNounPhrase()
	// 	}
	// }
	if randBoolChance(chance) {
		if randBoolChance(25) {
			loc = randNounPhrase()
			locPrep = "by"
		} else {
			loc = randLocationPhrase()
			locPrep = "at"
		}
		chance /= 2
	}
	if randBoolChance(chance) {
		if randBoolChance(50) {
			for ok := true; ok; ok = instr == nom || instr == acc || instr == dat {
				instr = randNounPhrase()
			}
		} else {
			instr = randItemPhrase()
		}
	}
	if randBoolChance(25) {
		adverb = choiceString(adverbs)
	}
	return verbPhrase{verb, nom, acc, dat, instr, loc, locPrep, adverb}
}

func randVerbPhraseFromContent(localActors, localItems, localLocations []nounPhrase) verbPhrase {
	var verb, locPrep, adverb string
	var nom, acc, dat, instr, loc nounPhrase
	var chance int = 50

	nom = choice(localActors)
	if randBoolChance(50) {
		verb = choiceString(transVerbs)
		for ok := true; ok; ok = (acc == nom) {
			acc = choice(localActors)
		}
	} else {
		verb = choiceString(intransVerbs)
	}
	// if randBoolChance(chance) {
	// 	for ok := true; ok; ok = dat == nom || dat == acc {
	// 		dat = randNounPhrase()
	// 	}
	// }
	if randBoolChance(chance) {
		if randBoolChance(25) {
			loc = choice(localActors)
			locPrep = "by"
		} else {
			loc = choice(localLocations)
			locPrep = "at"
		}
		chance /= 2
	}
	if randBoolChance(chance) {
		if randBoolChance(25) {
			for ok := true; ok; ok = instr == nom || instr == acc || instr == dat {
				instr = choice(localActors)
			}
		} else {
			instr = choice(localItems)
		}
	}
	if randBoolChance(25) {
		adverb = choiceString(adverbs)
	}
	return verbPhrase{verb, nom, acc, dat, instr, loc, locPrep, adverb}
}

func randVerbPhrase() verbPhrase {
	var actor nounPhrase = randNounPhrase()
	return randActorVerbPhrase(actor)
}
