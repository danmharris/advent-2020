package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func StringInRange(s string, min int, max int) bool {
	i, err := strconv.Atoi(s)
	return err == nil && i >= min && i <= max
}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (pp *Passport) HasRequiredFields() bool {
	return pp.byr != "" &&
		pp.iyr != "" &&
		pp.eyr != "" &&
		pp.hgt != "" &&
		pp.hcl != "" &&
		pp.ecl != "" &&
		pp.pid != ""
}

func (pp *Passport) HasValidFields() bool {
	// byr
	if !StringInRange(pp.byr, 1920, 2002) {
		return false
	}

	//iyr
	if !StringInRange(pp.iyr, 2010, 2020) {
		return false
	}

	// eyr
	if !StringInRange(pp.eyr, 2020, 2030) {
		return false
	}

	// hgt
	re := regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
	groups := re.FindStringSubmatch(pp.hgt)
	if groups == nil {
		return false
	}
	if groups[2] == "cm" && !StringInRange(groups[1], 150, 193) {
		return false
	}
	if groups[2] == "in" && !StringInRange(groups[1], 59, 76) {
		return false
	}

	// hcl
	if m, _ := regexp.MatchString(`^#[\da-f]{6}$`, pp.hcl); !m {
		return false
	}

	// ecl
	if m, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, pp.ecl); !m {
		return false
	}

	// pid
	if m, _ := regexp.MatchString(`^\d{9}$`, pp.pid); !m {
		return false
	}

	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var passports []Passport
	current := Passport{}

	for sc.Scan() {
		if sc.Text() == "" {
			passports = append(passports, current)
			current = Passport{}
			continue
		}

		fields := strings.Split(sc.Text(), " ")

		for _, f := range fields {
			kv := strings.Split(f, ":")
			var k *string
			switch kv[0] {
			case "byr":
				k = &current.byr
			case "iyr":
				k = &current.iyr
			case "eyr":
				k = &current.eyr
			case "hgt":
				k = &current.hgt
			case "hcl":
				k = &current.hcl
			case "ecl":
				k = &current.ecl
			case "pid":
				k = &current.pid
			case "cid":
				k = &current.cid
			default:
				panic("Unknown field")
			}

			*k = kv[1]
		}
	}
	passports = append(passports, current)

	present := 0
	valid := 0
	for _, pp := range passports {
		if pp.HasRequiredFields() {
			present++
		}
		if pp.HasValidFields() {
			valid++
		}
	}

	fmt.Printf("Number of present: %d\n", present)
	fmt.Printf("Number of valid: %d\n", valid)
}
