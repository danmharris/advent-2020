package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func IsSum(consider []int, toFind int) bool {
	for m := 0; m < len(consider); m++ {
		for n := 0; n < len(consider); n++ {
			if consider[m]+consider[n] == toFind {
				return true
			}
		}
	}

	return false
}

func FindInvalid(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		consider := numbers[i-preamble : i]

		if !IsSum(consider, numbers[i]) {
			return numbers[i]
		}
	}

	return 0
}

func SumSlice(numbers []int) int {
	var total int
	for _, i := range numbers {
		total += i
	}
	return total
}

func FindContiguousAdd(numbers []int, target int) []int {
	for size := 2; size < len(numbers); size++ {
		for i := 0; i <= len(numbers)-size; i++ {
			if consider := numbers[i : i+size]; SumSlice(consider) == target {
				return consider
			}
		}
	}

	return nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var numbers []int

	for sc.Scan() {
		line := sc.Text()

		n, _ := strconv.Atoi(line)

		numbers = append(numbers, n)
	}

	invalid := FindInvalid(numbers, 25)
	fmt.Printf("%d\n", invalid)

	contiguous := FindContiguousAdd(numbers, invalid)
	sort.Ints(contiguous)
	weakness := contiguous[0] + contiguous[len(contiguous)-1]

	fmt.Printf("%d\n", weakness)
}
