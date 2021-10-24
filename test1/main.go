package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		// Seeding with the same value results in the same random sequence each run.
		// For different numbers, seed with a different value, such as
		// time.Now().UnixNano(), which yields a constantly-changing number.
		rand.Seed(1)
		answers := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes definitely",
			"You may rely on it",
			"As I see it yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}
		fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	*/

	max := 10
	u, _ := time.ParseDuration(("2us"))
	fmt.Printf("One microsencond is %d nanoseconds,\n", u.Nanoseconds())

	t := time.Now().UnixNano()

	rand.Seed(t)

	fmt.Println(rand.Intn(max))

}
