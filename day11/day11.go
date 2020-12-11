package main

import (
	"bufio"
	"fmt"
	"os"
)

type seat int
type directionFunc func(int, int) (int, int)
type countFunc func(*[][]seat, int, int) int

const (
	Floor seat = iota
	Occupied
	Unoccupied
)

var mapping = map[rune]seat{
	'.': Floor,
	'#': Occupied,
	'L': Unoccupied,
}

var directionFuncs = [8]directionFunc{
	func(x int, y int) (int, int) { return x, y - 1 },
	func(x int, y int) (int, int) { return x + 1, y - 1 },
	func(x int, y int) (int, int) { return x + 1, y },
	func(x int, y int) (int, int) { return x + 1, y + 1 },
	func(x int, y int) (int, int) { return x, y + 1 },
	func(x int, y int) (int, int) { return x - 1, y + 1 },
	func(x int, y int) (int, int) { return x - 1, y },
	func(x int, y int) (int, int) { return x - 1, y - 1 },
}

// Part 2: Count the first seats seen in each direction
func CountAdjacentFirst(seats *[][]seat, xI int, yI int) int {
	count := 0
	s := *seats

	for _, f := range directionFuncs {
		x, y := xI, yI
		for true {
			x, y = f(x, y)
			if y < 0 || y >= len(s) || x < 0 || x >= len(s[y]) {
				break
			}
			if s[y][x] != Floor {
				if s[y][x] == Occupied {
					count++
				}
				break
			}
		}
	}

	return count
}

// Part 1: Count the immediately adjacent seats in each direction
func CountAdjacentImmediate(seats *[][]seat, xI int, yI int) int {
	count := 0
	s := *seats

	for _, f := range directionFuncs {
		x, y := f(xI, yI)

		if y >= 0 && y < len(s) && x >= 0 && x < len(s[y]) && s[y][x] == Occupied {
			count++
		}
	}

	return count
}

// Count how many seats in the grid are occupied
func CountOccupied(seats [][]seat) int {
	var total int
	for _, y := range seats {
		for _, x := range y {
			if x == Occupied {
				total++
			}
		}
	}
	return total
}

// Perform one iteration of the rules
func Iterate(seats [][]seat, f countFunc, occThreshold int) ([][]seat, bool) {
	var (
		newSeats [][]seat
		change   bool
	)

	for yI, y := range seats {
		var row []seat

		for xI, x := range y {
			newSeat := x

			if x == Occupied {
				if adj := f(&seats, xI, yI); adj >= occThreshold {
					newSeat = Unoccupied
					change = true
				}
			} else if x == Unoccupied {
				if adj := f(&seats, xI, yI); adj == 0 {
					newSeat = Occupied
					change = true
				}
			}

			row = append(row, newSeat)
		}

		newSeats = append(newSeats, row)
	}

	return newSeats, change
}

// Convert file to grid
func ParseSeats(sc *bufio.Scanner) [][]seat {
	var seats [][]seat

	for sc.Scan() {
		var row []seat
		l := sc.Text()

		for _, s := range l {
			row = append(row, mapping[s])
		}
		seats = append(seats, row)
	}

	return seats
}

// Keep running iterations until they stabilise
func Simulate(seats [][]seat, f countFunc, occThreshold int) int {
	change := true

	for change {
		seats, change = Iterate(seats, f, occThreshold)
	}

	return CountOccupied(seats)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	seats := ParseSeats(sc)

	fmt.Printf("occupied (part 1): %d\n", Simulate(seats, CountAdjacentImmediate, 4))
	fmt.Printf("occupied (part 2): %d\n", Simulate(seats, CountAdjacentFirst, 5))
}
