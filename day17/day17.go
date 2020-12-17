package main

import (
	"bufio"
	"fmt"
	"os"
)

type State int

const (
	Inactive State = iota
	Active
)

type Position struct {
	x int
	y int
	z int
	w int
}

type Grid map[Position]State

func (g Grid) CountAdjacent(pos Position) int {
	var adj int
	for zI := pos.z - 1; zI <= pos.z+1; zI++ {
		for yI := pos.y - 1; yI <= pos.y+1; yI++ {
			for xI := pos.x - 1; xI <= pos.x+1; xI++ {
				for wI := pos.w - 1; wI <= pos.w+1; wI++ {
					p := Position{
						x: xI,
						y: yI,
						z: zI,
						w: wI,
					}

					if s, ok := g[p]; ok && p != pos && s == Active {
						adj++
					}
				}
			}
		}
	}

	return adj
}

func (g Grid) CountActive() int {
	var total int
	for _, s := range g {
		if s == Active {
			total++
		}
	}
	return total
}

func (grid *Grid) Pad(withW bool) {
	var (
		minX int = 1000
		maxX int = -1000
		minY int = 1000
		maxY int = -1000
		minZ int = 1000
		maxZ int = -1000
		minW int = 1000
		maxW int = -1000
	)
	for p, _ := range *grid {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.z < minZ {
			minZ = p.z
		}
		if p.z > maxZ {
			maxZ = p.z
		}
		if p.w < minW {
			minW = p.w
		}
		if p.w > maxW {
			maxW = p.w
		}
	}

	for z := minZ - 1; z <= maxZ+1; z++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				for w := minW - 1; w <= maxW+1; w++ {
					p := Position{
						x: x,
						y: y,
						z: z,
					}
					if withW {
						p.w = w
					}
					if _, ok := (*grid)[p]; !ok {
						(*grid)[p] = Inactive
					}
				}
			}
		}
	}
}

func (grid Grid) Iterate() Grid {
	newGrid := make(Grid)
	grid.Pad(true) // Change to false for part 1

	for p, s := range grid {
		adj := grid.CountAdjacent(p)
		newS := s
		switch s {
		case Inactive:
			if adj == 3 {
				newS = Active
			}
		case Active:
			if !(adj == 2 || adj == 3) {
				newS = Inactive
			}
		}

		newGrid[p] = newS
	}

	return newGrid
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	grid := make(Grid)
	var y int
	for sc.Scan() {
		for x, char := range sc.Text() {
			var s State
			switch char {
			case '.':
				s = Inactive
			case '#':
				s = Active
			}

			grid[Position{z: 0, y: y, x: x, w: 0}] = s
		}
		y--
	}

	for i := 0; i < 6; i++ {
		grid = grid.Iterate()
	}

	total := grid.CountActive()
	fmt.Printf("%d\n", total)
}
