package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetCountsAny(answers *map[rune]int, numPeople int) int {
	return len(*answers)
}

func GetCountsAll(answers *map[rune]int, numPeople int) int {
	var sum int
	for _, c := range *answers {
		if c == numPeople {
			sum++
		}
	}
	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var (
		sumPartOne int
		sumPartTwo int
		numPeople  int
		answers    map[rune]int = make(map[rune]int)
	)

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			sumPartOne += GetCountsAny(&answers, numPeople)
			sumPartTwo += GetCountsAll(&answers, numPeople)
			answers = make(map[rune]int)
			numPeople = 0
			continue
		}

		numPeople++

		for _, c := range line {
			answers[c]++
		}
	}
	sumPartOne += GetCountsAny(&answers, numPeople)
	sumPartTwo += GetCountsAll(&answers, numPeople)

	fmt.Printf("Total (any): %d\n", sumPartOne)
	fmt.Printf("Total (all): %d\n", sumPartTwo)
}
