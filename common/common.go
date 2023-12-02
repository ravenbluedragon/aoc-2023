package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// const for output colour codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

// ChoosePart runs the appropriate function based on the provided part number
func ChoosePart(day int, part1 func(string) any, part2 func(string) any) {
	filename := fmt.Sprintf("data/%02d.txt", day)

	if len(os.Args) < 2 {
		log.Fatal("Missing argument: part1 or part2")
	}

	var value any
	part := os.Args[1]

	switch part {
	case "1":
		value = part1(filename)
	case "2":
		value = part2(filename)
	default:
		log.Fatalf("Expected part to be 1 or 2, was %s", part)
	}

	formatString := fmt.Sprintf("%sDay %%d%s %sPart %%s%s: %s%%v%s\n", Green, Reset, Purple, Reset, Bold+Yellow, Reset)
	log.Printf(formatString, day, part, value)
}

// LoadData reads lines from a file and returns them as a slice of strings
func LoadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
