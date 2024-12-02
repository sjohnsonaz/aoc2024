package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sjohnsonaz/aoc2024/day01/solution"
)

func main() {
	name := getFileName()
	fmt.Printf("analyzing %v...\n", name)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	sd, err := solution.NewSafeData(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Safe: %v, Dampened: %v\n", sd.Safe, sd.SafeDampened)
}

func getFileName() string {
	if len(os.Args) < 2 {
		return "data.csv"
	}
	return os.Args[1]
}
