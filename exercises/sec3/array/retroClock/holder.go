package main

type Holder [5]string

var (
	Zero = Holder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	One = Holder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	Two = Holder{
		"███",
		"  █",
		" █ ",
		"█  ",
		"███",
	}

	Three = Holder{
		"███",
		"  █",
		" ██",
		"  █",
		"███",
	}

	Four = Holder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	Five = Holder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	Six = Holder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	Seven = Holder{
		"███",
		"  █",
		" █ ",
		"█  ",
		"█  ",
	}

	Eight = Holder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	Nine = Holder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	Sep = Holder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
)

var Numbers = [...]Holder{
	Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Sep,
}
