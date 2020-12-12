package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Ship struct {
	x   int
	y   int
	d   dir
	way *Ship
}

// Part 1
func (s *Ship) Move(d rune, n int) {
	if d == 'F' {
		d = s.d.ToRune()
	}
	switch d {
	case 'E':
		s.x += n
	case 'S':
		s.y -= n
	case 'W':
		s.x -= n
	case 'N':
		s.y += n
	case 'L':
		s.d = s.d.Rotate(-n)
	case 'R':
		s.d = s.d.Rotate(n)
	}
}

// Part 2
func (s *Ship) MoveWithWaypoint(d rune, n int) {
	switch d {
	case 'L':
		s.way.RotateRelToZero(-n)
	case 'R':
		s.way.RotateRelToZero(n)
	case 'F':
		s.x += n * s.way.x
		s.y += n * s.way.y
	default:
		s.way.Move(d, n)
	}
}

func (s *Ship) RotateRelToZero(a int) {
	if a > 0 {
		for i := 0; i < a; i += 90 {
			s.x, s.y = s.y, -s.x
		}
	} else {
		for i := 0; i > a; i -= 90 {
			s.x, s.y = -s.y, s.x
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

	s := Ship{}
	s2 := Ship{
		way: &Ship{
			x: 10,
			y: 1,
		},
	}

	for sc.Scan() {
		m := re.FindStringSubmatch(sc.Text())
		d := rune(m[1][0])
		n, err := strconv.Atoi(m[2])

		if err != nil {
			panic(err)
		}

		s.Move(d, n)
		s2.MoveWithWaypoint(d, n)
	}

	x := int(math.Abs(float64(s.x)))
	y := int(math.Abs(float64(s.y)))
	x2 := int(math.Abs(float64(s2.x)))
	y2 := int(math.Abs(float64(s2.y)))

	fmt.Printf("%d\t%d\t%d\n", x, y, x+y)
	fmt.Printf("%d\t%d\t%d\n", x2, y2, x2+y2)
}
