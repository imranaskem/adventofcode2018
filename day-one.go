package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOne(data []string) {
	lines := data

	freqs := make(map[int]int)

	answer := 0
	dupe := 0
	dupeFound := false
	loops := 0

	for !dupeFound {
		fmt.Println(loops)

		for _, i := range lines {
			isPlus := strings.Contains(i, "+")
			numString := strings.TrimLeft(i, "+-")
			num, _ := strconv.Atoi(numString)

			if isPlus {
				answer += num
			} else {
				answer -= num
			}

			freqs[answer]++

			if freqs[answer] >= 2 {
				dupe = answer
				dupeFound = true
				break
			}
		}

		loops++
	}

	fmt.Println(answer)
	fmt.Println(dupe)
}
