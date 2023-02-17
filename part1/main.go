package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Create a program that will read in a quiz provided via a CSV file
// and will then give the quiz to a user keeping track of how many questions they get right
// and how many they get incorrect. Regardless of whether the answer is correct or wrong
// the next question should be asked immediately afterwards.

// The CSV file should default to `problems.csv` (example shown below),
// but the user should be able to customize the filename via a flag.

// The first column of the CSV is a question and the second column in the same row is the answer

// You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

// At the end of the quiz the program should output the total number of questions correct and how many questions there were in total.
// Questions given invalid answers are considered incorrect.

// **NOTE:** *CSV files may have questions with commas in them.
// Eg: `"what 2+2, sir?",4` is a valid row in a CSV.
// I suggest you look into the CSV package in Go and don't try to write your own CSV parser.*

func main() {
	fmt.Println("Welcome to this super fancy quiz game")
	readCSV()
}

func readCSV() {
	quizFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(quizFile)

	records, _ := reader.ReadAll()

	fmt.Println(records)
}
