package main

type dir int

const (
	East dir = iota
	South
	West
	North
)

func (d dir) ToRune() rune {
	switch d {
	case East:
		return 'E'
	case South:
		return 'S'
	case West:
		return 'W'
	case North:
		return 'N'
	}

	panic("Unknown")
}

func (d dir) Reverse() dir {
	switch d {
	case East:
		return West
	case South:
		return North
	case West:
		return East
	case North:
		return South
	}

	panic("Unknown")
}

func (d dir) Rotate(n int) dir {
	nd := int(float64(d) + float64(n/90))
	if nd < 0 {
		nd += 4
	}
	return dir(nd % 4)
}
