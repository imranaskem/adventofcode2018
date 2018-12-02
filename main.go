package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)

	freqs := make(map[int]int)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

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
