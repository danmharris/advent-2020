package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())

		if err == nil {
			numbers = append(numbers, n)
		}
	}

	sort.Ints(numbers)

	mid := sort.Search(len(numbers), func(i int) bool { return numbers[i] >= 1010 })
	lowerHalf := numbers[:mid]

	quarter := sort.Search(len(lowerHalf), func(i int) bool { return numbers[i] >= 505 })
	lowerQuarter := lowerHalf[:quarter]

	for _, x := range numbers {
		for _, y := range lowerHalf {
			if x+y == 2020 {
				fmt.Printf("%d * %d = %d\n", x, y, x*y)
			}
			for _, z := range lowerQuarter {
				if x+y+z == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", x, y, z, x*y*z)
				}
			}
		}
	}
}
