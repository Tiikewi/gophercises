package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	// get filename
	var fileName = flag.String("file", "problems.csv", "quiz file name.")
	var seconds = flag.Int("time", 30, "Amount of time to take the quiz.")
	flag.Parse()
	problems := readCSV(*fileName)

	var correct int = 0
	var incorrect int = 0

	// start timer
	go bgTimer(*seconds, &correct)

	for _, problem := range problems {
		var answer string
		fmt.Printf("%v:\n", problem.Question)
		_, err := fmt.Scan(&answer)
		if err != nil {
			log.Fatal("Error when reading user answer: ", err)
		}
		if answer == problem.Answer {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("%v out of %v correct.\n", correct, correct+incorrect)
}

func bgTimer(seconds int, correct *int) {
	// create timer
	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	<-timer.C
	fmt.Printf("You got %v correct answers before time ended!\n", *correct)
	os.Exit(1)
}
