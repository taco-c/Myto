package main

var (
	actors = []string{
		// Animals
		"sheep",     //
		"ram",       //
		"wether",    // Castrated male sheep or goat
		"ewe",       //
		"lamb",      //
		"goat",      //
		"billy",     // Male goat
		"nanny",     // Female goat
		"kid",       // Young goat
		"cock",      // Male chicken
		"hen",       //
		"chick",     // Young chicken
		"cockerel",  // Male young chicken
		"pullet",    // Female young chicken
		"dog",       //
		"stud",      // Male dog
		"bitch",     // Female dog
		"puppy",     //
		"cat",       //
		"tomcat",    // Male cat
		"gibcat",    // Castrated male cat
		"mollycat",  // Female cat
		"kitten",    //
		"cow",       //
		"bull",      // Uncastrated male cattle
		"ox",        // Castrated male cattle
		"calf",      //
		"horse",     //
		"stallion",  // Uncastrated male horse
		"mare",      // Female horse
		"filly",     // Female young horse
		"donkey",    //
		"jack",      // Male donkey
		"jenny",     // Female donkey
		"foal",      // Young donkey / young horse
		"dove",      //
		"goose",     //
		"gander",    // Male goose
		"gosling",   //
		"bear boar", //
		"bear sow",  //
		"bear cub",  //
		"pig",       //
		"boar",      //
		"sow pig",   //
		"piglet",    //
		"bird",      //
		"fish",      //
		"bat",       //
		"seagull",   //
		"eagle",     //
		"hawk",      //
		"tiercel",   // Male hawk
		"hawk hen",  // Female hawk
		"falcon",    //
		// Humanoids
		"man",
		"woman",
		"elf",
		"dwarf",
		"gnome",
		// Occupations
		"god",
		"godess",
		"king",
		"queen",
		"duke",
		"duchess",
		"count",
		"countess",
		"baron",
		"baroness",
		"lord",
		"lady",
		"advisor",
		"prophet",
		"carpenter",
		"dancer",
		"singer",
		"actor",
		"farmer",
		"trader",
		"thief",
		"guardian",
		"guard",
	}
	actoractors = [][]string{
		{"a", "sheep", "sheep"},
		{"a", "ram", "rams"},
		{"a", "wether", "wethers"},
		{"a", "ewe", "ewes"},
		{"a", "lamb", "lambs"},
	}

	items = []string{
		"apple",
		"candle",
		"fork",
		"wool",
		"apple",
		"egg",
		"bed",
		"tree",
		"axe",
		"cup",
		"vase",
		"stone",
		"fruit",
		"shirt",
		"bread",
		"potato",
		"tomato",
		"sword",
		"bucket",
		"bowl",
		"book",
		"knife",
	}

	// In the beginning there is beautiful Falcon.
	// The beautiful Falcon thinks of delight near the beautiful Falcon.
	// The beautiful Falcon is the symbol of friendship.

	// In the beginning there is plain Seagull of Joy.
	// The plain Seagull of Joy thinks of fascination.
	// The plain Seagull of Joy joins itself.
	// Then the delightful Potato comes.
	// Then the delightful Potato sings, making the delightful Potato calm.
	// The plain Seagull of Joy is the symbol of luxury.
	// The delightful Potato is the symbol of pain.

	abstractStuff = []string{
		"love",
		"tax",
		"curiousity",
		"wealth",
		"adoration",
		"amazement",
		"anger",
		"clarity",
		"delight",
		"despair",
		"disbelief",
		"belief",
		"fascination",
		"friendship",
		"grief",
		"happiness",
		"hate",
		"joy",
		"misery",
		"pain",
		"pleasure",
		"pride",
		"luxury",
		"power",
		"relief",
		"sorrow",
		"strenght",
		"surprise",
		"tiredness",
		"sleep",
		"worry",
		"faith",
		"fate",
		"pureness",
		"temperance",
		"wisdom",
		"justice",
		"courage",
		"prudence",
		"fortitude",
		"chastity",
		"charity",
		"diligence",
		"patience",
		"kindness",
		"humility",
		"lust",
		"gluttony",
		"greed",
		"sloth",
		"wrath",
		"envy",
		"pride",
		"fear",
		"risk",
		"cowardice",
	}

	adjectives = []string{
		"red",
		"blue",
		"yellow",
		"purple",
		"green",
		"orange",
		"big",
		"small",
		"nice",
		"good",
		"bad",
		"rude",
		"white",
		"black",
		"grey",
		"light",
		"tall",
		"short",
		"beautiful",
		"fit",
		"handsome",
		"long",
		"muscular",
		"plain",
		"chubby",
		"plump",
		"skinny",
		"ugly",
		"unsightly",
		"unkempt",
		"ambitious",
		"brave",
		"calm",
		"delightful",
		"eager",
		"faithful",
		"gentle",
		"zealous",
		"bewildered",
		"fierce",
		"jealous",
		"pitiful",
		"repulsive",
		"nervous",
		"stubborn",
		"hallow",
		"ancient",
		"early",
		"old",
		"young",
		"pure",
		"decent",
		"temperant",
		"wise",
		"just",
		"couragous",
		"prude",
		"chaste",
		"diligent",
		"patient",
		"kind",
		"humble",
		"lusty",
		"gluttonous",
		"greedy",
		"lazy",
		"wrathful",
		"angry",
		"envious",
		"proud",
		"cowardly",
	}

	adverbs = []string{
		"quickly",
		"slowly",
		"tirelessly",
		"decently",
		"repeatedly",
		"happily",
		"nicely",
		"badly",
	}

	intransVerbs = []string{
		"sits",
		"stands",
		"talks",
		"walks",
		"sleeps",
		"sings",
		"dances",
		"wiggles",
		"works",
		"yells",
		"agrees",
		"appears",
		"arrives",
		"becomes",
		"belongs",
		"collapses",
		"consists",
		"coughes",
		"cries",
		"dies",
		"disappears",
		"emerges",
		"exists",
		"explodes",
		"fades",
		"falls",
		"fasts",
		"floats",
		"flies",
		"gallops",
		"goes",
		"grows",
		"happens",
		"inquires",
		"jumps",
		"kneels",
		"knocks",
		"lasts",
		"laughs",
		"leads",
		"leans",
		"lies",
		"listens",
		"lives",
		"looks",
		"marches",
		"mourns",
		"moves",
		"panics",
		"pauses",
		"peeps",
		"poses",
		"pounces",
		"pouts",
		"prays",
		"reclines",
		"relaxes",
		"remains",
		"revolts",
		"rises",
		"rolls",
		"runs",
		"screams",
		"shakes",
		"shouts",
		"sighs",
		"sits",
		"skips",
		"sleeps",
		"slides",
		"smells",
		"smiles",
		"snarls",
		"sneezes",
		"soaks",
		"spins",
		"spits",
		"sprints",
		"squeaks",
		"staggers",
		"stays",
		"swims",
		"swings",
		"twists",
		"vanishes",
		"vomits",
		"wades",
		"waits",
		"walks",
		"wanders",
		"waves",
		"whirls",
	}

	transVerbs = []string{
		"devours",
		"accepts",
		"dazzles",
		"deceives",
		"beats",
		"discovers",
		"blesses",
		"avoids",
		"captures",
		"carries",
		"entertains",
		"follows",
		"cleans",
		"freezes",
		"warms",
		"hugs",
		"frightens",
		"forgives",
		"converts",
		"comforts",
		"grabs",
		"knocks",
		"kisses",
		"honours",
		"helps",
		"impresses",
		"loves",
		"likes",
		"adores",
		"hates",
		"despices",
		"inpects",
		"marries",
		"maintains",
		"mocks",
		"interrupts",
		"joins",
		"notices",
		"judges",
		"kicks",
		"punches",
		"sees",
		"pleases",
		"teaches",
		"tightens",
		"pierces",
		"prepares",
		"turns",
		"protects",
		"understands",
		"questions",
		"lifts",
		"uses",
		"reminds",
		"wants",
		"washes",
		"watches",
		"warns",
		"scolds",
		"satisfies",
		"slaps",
		"wears",
		"smells",
		"softens",
		"spreads",
	}

	prepositions = []string{
		"on",
		"under",
		"near",
		"at",
		"to",
		"with",
		"without",
	}

	locations = []string{
		"mountain",
		"valley",
		"river",
		"hill",
		"city",
		"town",
		"temple",
		"home",
		"marketplace",
		"bridge",
		"thone",
		"hall",
		"lake",
		"island",
		"peninsula",
		"village",
		"castle",
		"tower",
		"theatre",
		"festival",
		"chamber",
		"square",
	}
)
