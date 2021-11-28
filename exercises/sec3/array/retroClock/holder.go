package main

type Holder [5]string

var (
	zero = Holder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = Holder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = Holder{
		"███",
		"  █",
		" █ ",
		"█  ",
		"███",
	}

	three = Holder{
		"███",
		"  █",
		" ██",
		"  █",
		"███",
	}

	four = Holder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = Holder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = Holder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven = Holder{
		"███",
		"  █",
		" █ ",
		"█  ",
		"█  ",
	}

	eight = Holder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = Holder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	sep = Holder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
)

var Numbers = [...]Holder{
	zero, one, two, three, four, five, six, seven, eight, nine, sep,
}
