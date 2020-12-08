package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Cmd int

const (
	Nop Cmd = iota
	Acc
	Jmp
)

type Instruction struct {
	cmd Cmd
	val int
}

func ParseCode(sc *bufio.Scanner) []*Instruction {
	var code []*Instruction

	re := regexp.MustCompile(`(\w+) ((?:\+|-)\d+)`)

	for sc.Scan() {
		line := re.FindStringSubmatch(sc.Text())

		var cmd Cmd
		switch line[1] {
		case "nop":
			cmd = Nop
		case "acc":
			cmd = Acc
		case "jmp":
			cmd = Jmp
		}
		val, _ := strconv.Atoi(line[2])

		code = append(code, &Instruction{
			cmd: cmd,
			val: val,
		})
	}

	return code
}

func FindLoop(code []*Instruction) (int, bool) {
	var (
		acc  int
		line int
		ins  *Instruction
		seen map[int]bool = make(map[int]bool)
	)

	for true {
		if seen[line] {
			return acc, true
		}
		if line >= len(code) {
			return acc, false
		}

		seen[line] = true
		ins = code[line]

		switch ins.cmd {
		case Nop:
			line++
		case Acc:
			acc += ins.val
			line++
		case Jmp:
			line += ins.val
		}
	}

	return 0, false
}

func FixLoop(code []*Instruction) int {
	var tmpCmd Cmd
	for _, i := range code {
		if i.cmd == Acc {
			continue
		}

		if i.cmd == Nop {
			tmpCmd = Nop
			i.cmd = Jmp
		} else if i.cmd == Jmp {
			tmpCmd = Jmp
			i.cmd = Nop
		}

		if acc, loop := FindLoop(code); !loop {
			return acc
		}
		i.cmd = tmpCmd
	}

	return 0
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	code := ParseCode(sc)
	acc, _ := FindLoop(code)

	fmt.Printf("With loop: %d\n", acc)
	fmt.Printf("With fixed loop: %d\n", FixLoop(code))
}
