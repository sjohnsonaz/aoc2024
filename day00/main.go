package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sjohnsonaz/aoc2024/day00/solution"
)

func main() {
	name := getFileName()
	fmt.Printf("analyzing %v...\n", name)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	la, err := solution.NewListAnalysis(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Distance: %v, Similarity: %v\n", la.Distance, la.Similarity)
}

func getFileName() string {
	if len(os.Args) < 2 {
		return "data.csv"
	}
	return os.Args[1]
}
