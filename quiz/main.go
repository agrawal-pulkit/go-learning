package main

import (
	"encoding/csv"
	"flag"
	"strings"
	"time"

	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the form of 'question, answer'.")
	timeLimit := flag.Int("limit", 30, "the limit for the quiz in seconds.")
	flag.Parse()
	_ = csvFileName
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprint("Failed in open csv file: %S\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to pass the provided csv file."))
	}
	problem := parseLines(lines)
	// fmt.Println(problem)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problem {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problem))
			return
		case answer := <-answerChannel:
			if answer == p.a {
				correct++
			}
		}
	}

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
