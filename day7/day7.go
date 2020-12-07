package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagRule struct {
	count int
	bag   string
}

var rules map[string][]*BagRule

func ContainsColour(colour string, rule string) bool {
	if rules[rule] == nil {
		return false
	}

	for _, r := range rules[rule] {
		if r.bag == colour || ContainsColour(colour, r.bag) {
			return true
		}
	}

	return false
}

func TotalBags(rule string) int {
	if rules[rule] == nil {
		return 0
	}

	sum := 0
	for _, r := range rules[rule] {
		sum += r.count + (r.count * TotalBags(r.bag))
	}

	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	rules = make(map[string][]*BagRule)
	re := regexp.MustCompile(`^(.*) bags contain (.*)[.]$`)
	ruleRe := regexp.MustCompile(`(\d)+ (.*) bags?`)

	for sc.Scan() {
		line := sc.Text()

		var contents []*BagRule
		m := re.FindStringSubmatch(line)
		if m[2] == "no other bags" {
			continue
		}

		for _, b := range strings.Split(m[2], ",") {
			rm := ruleRe.FindStringSubmatch(b)
			count, _ := strconv.Atoi(rm[1])

			rule := BagRule{
				count: count,
				bag:   rm[2],
			}

			contents = append(contents, &rule)
		}

		rules[m[1]] = contents
	}

	matching := 0
	for r, _ := range rules {
		if ContainsColour("shiny gold", r) {
			matching++
		}
	}

	fmt.Printf("%d\n", matching)
	fmt.Printf("%d\n", TotalBags("shiny gold"))
}
