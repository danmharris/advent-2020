package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Password struct {
	min      int
	max      int
	char     byte
	password string
}

func (pw *Password) IsValidPartOne() bool {
	freq := make(map[byte]int)

	for i := 0; i < len(pw.password); i++ {
		c := pw.password[i]
		_, ok := freq[c]

		if !ok {
			freq[c] = 1
		} else {
			freq[c]++
		}
	}

	if freq[pw.char] >= pw.min && freq[pw.char] <= pw.max {
		return true
	}
	return false
}

func (pw *Password) IsValidPartTwo() bool {
	first := (pw.password[pw.min-1] == pw.char)
	second := (pw.password[pw.max-1] == pw.char)
	return first != second
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	validPartOne := 0
	validPartTwo := 0
	for scanner.Scan() {
		line := regexp.MustCompile(`^(\d+)-(\d+) (\w): (.*)$`)

		parts := line.FindStringSubmatch(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		min, errMin := strconv.Atoi(parts[1])
		max, errMax := strconv.Atoi(parts[2])
		if errMin != nil || errMax != nil {
			panic("invalid characters")
		}

		newPassword := Password{
			min:      min,
			max:      max,
			char:     parts[3][0],
			password: parts[4],
		}

		if newPassword.IsValidPartOne() {
			validPartOne++
		}
		if newPassword.IsValidPartTwo() {
			validPartTwo++
		}
	}

	fmt.Printf("%d\n", validPartOne)
	fmt.Printf("%d\n", validPartTwo)
}
