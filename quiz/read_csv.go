package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCSV(fileName string) []Problem {

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error when closing file named %v: %v", fileName, err)
		}
	}(file)

	// Read csv
	csvReader := csv.NewReader(file)

	questions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error when reading CSV file: ", err)
	}

	return createProblem(questions)

}

func createProblem(questions [][]string) []Problem {
	var problems []Problem

	for _, line := range questions {
		var problem Problem

		for j, field := range line {
			if j == 0 {
				problem.Question = field

			} else {
				problem.Answer = field
			}

		}
		problems = append(problems, problem)
	}

	return problems
}
