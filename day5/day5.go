package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var seatIds []int

	for sc.Scan() {
		pass := sc.Text()

		var (
			maxRow float64 = 127
			minRow float64 = 0
			maxCol float64 = 7
			minCol float64 = 0
		)

		for _, c := range pass {
			midRow := ((maxRow-minRow)/2 + minRow)
			midCol := ((maxCol-minCol)/2 + minCol)
			switch c {
			case 'F':
				maxRow = math.Floor(midRow)
			case 'B':
				minRow = math.Ceil(midRow)
			case 'L':
				maxCol = math.Floor(midCol)
			case 'R':
				minCol = math.Ceil(midCol)
			}
		}

		if minRow != maxRow || minCol != maxCol {
			panic("Couldn't narrow down to 1 seat!")
		}

		seatId := int(maxRow)*8 + int(maxCol)
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	fmt.Printf("Max seat ID: %d\n", seatIds[len(seatIds)-1])

	currSeatId := seatIds[0] - 1
	for _, id := range seatIds {
		currSeatId++
		if currSeatId != id {
			fmt.Printf("Missing seat ID found: %d\n", currSeatId)
			break
		}
	}
}
