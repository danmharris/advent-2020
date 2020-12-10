package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Seq(n int) int {
	total := 1
	for i := n - 1; i > 0; i-- {
		total += i
	}
	return total
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var adapters []int

	for sc.Scan() {
		j, _ := strconv.Atoi(sc.Text())
		adapters = append(adapters, j)
	}

	sort.Ints(adapters)

	var (
		prev  int
		one   int
		three int = 1
	)

	for _, j := range adapters {
		if j-prev == 1 {
			one++
		} else if j-prev == 3 {
			three++
		}
		prev = j
	}

	device := prev + 3
	fmt.Printf("%d\n", one*three)

	adapters = append(adapters, device)
	combos := 1
	lastThree := 0

	for i := 0; i < len(adapters)-1; i++ {
		if adapters[i+1]-adapters[i] == 3 {
			diff := adapters[i] - lastThree
			combos *= Seq(diff)

			lastThree = adapters[i+1]
		}
	}

	fmt.Printf("%d\n", combos)
}
