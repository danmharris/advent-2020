package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	nums := make(map[int][]int)
	var (
		prev int
		turn int
	)

	sc.Scan()
	line := strings.Split(sc.Text(), ",")

	for _, n := range line {
		turn++

		num, _ := strconv.Atoi(n)
		nums[num] = []int{turn}
		prev = num
	}

	for turn < 30000000 {
		turn++

		var diff int
		if len(nums[prev]) > 1 {
			diff = nums[prev][len(nums[prev])-1] - nums[prev][len(nums[prev])-2]
		}
		nums[diff] = append(nums[diff], turn)
		prev = diff

		// Part 1
		if turn == 2020 {
			fmt.Printf("%d\n", prev)
		}
	}

	fmt.Printf("%d\n", prev)
}
