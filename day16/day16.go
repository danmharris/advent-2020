package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ticket []int

type Rule struct {
	name string
	aMin int
	aMax int
	bMin int
	bMax int
}

func NewRuleFromString(s string) *Rule {
	re := regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)
	m := re.FindStringSubmatch(s)

	aMin, _ := strconv.Atoi(m[2])
	aMax, _ := strconv.Atoi(m[3])
	bMin, _ := strconv.Atoi(m[4])
	bMax, _ := strconv.Atoi(m[5])

	return &Rule{
		name: m[1],
		aMin: aMin,
		aMax: aMax,
		bMin: bMin,
		bMax: bMax,
	}
}

func (r *Rule) IsValid(n int) bool {
	return (n >= r.aMin && n <= r.aMax) || (n >= r.bMin && n <= r.bMax)
}

func GetOnlyTrue(list []bool) (int, bool) {
	var (
		foundOne bool
		recent   int
	)
	for i, v := range list {
		if v {
			if foundOne {
				return 0, false
			}
			recent = i
			foundOne = true
		}
	}

	return recent, foundOne
}

func GetLowest(possibleFields *[][]bool) (int, int) {
	for ruleI, rule := range *possibleFields {
		if fieldI, found := GetOnlyTrue(rule); found {
			return ruleI, fieldI
		}
	}
	return -1, -1
}

func ClearLowest(possibleFields *[][]bool, r int, f int) {
	for ruleI, rule := range *possibleFields {
		for fieldI, _ := range rule {
			if ruleI == r || fieldI == f {
				rule[fieldI] = false
			}
		}
	}
}

func FillBoolArray(height int, width int) *[][]bool {
	table := make([][]bool, height)
	for y, _ := range table {
		row := make([]bool, width)
		for x, _ := range row {
			row[x] = true
		}
		table[y] = row
	}
	return &table
}

func main() {
	var (
		myTicket     Ticket
		validTickets []Ticket
		rules        []*Rule
		rate         int
	)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		if sc.Text() == "" {
			break
		}

		rules = append(rules, NewRuleFromString(sc.Text()))
	}

	for sc.Scan() {
		parts := strings.Split(sc.Text(), ",")
		if len(parts) == 1 {
			continue
		}

		var (
			t           Ticket
			ticketValid bool = true
			thisRate    int
		)
		for _, s := range parts {
			i, _ := strconv.Atoi(s)
			t = append(t, i)

			var isValid bool
			for _, r := range rules {
				if r.IsValid(i) {
					isValid = true
				}
			}

			if !isValid {
				thisRate += i
				ticketValid = false
			}
		}

		if myTicket == nil {
			myTicket = t
			continue
		}

		if !ticketValid {
			rate += thisRate
		} else {
			validTickets = append(validTickets, t)
		}
	}

	fmt.Printf("%d\n", rate)

	// rule -> field -> bool
	possibleFields := FillBoolArray(len(rules), len(myTicket))
	for _, t := range validTickets {
		for fieldI, fieldV := range t {
			for ruleI, rule := range rules {
				if !rule.IsValid(fieldV) {
					(*possibleFields)[ruleI][fieldI] = false
				}
			}
		}
	}

	var (
		ruleI   int
		fieldI  int
		product int = 1
	)

	re := regexp.MustCompile(`^departure`)
	for true {
		ruleI, fieldI = GetLowest(possibleFields)

		if ruleI < 0 && fieldI < 0 {
			break
		}

		if re.MatchString(rules[ruleI].name) {
			product *= myTicket[fieldI]
		}

		ClearLowest(possibleFields, ruleI, fieldI)
	}

	fmt.Printf("%d\n", product)
}
