package main

import (
	"fmt"
	"strings"
)

func dayTwo(filename string) {
	linesOfText := loadFromTextFile(filename)

	containsTwoLetters := 0
	containsThreeLetters := 0

	for _, line := range linesOfText {

		doesTwo := false
		doesThree := false

		for _, letter := range line {
			numOfRunes := strings.Count(line, string(letter))

			if numOfRunes == 2 {
				doesTwo = true
			}

			if numOfRunes == 3 {
				doesThree = true
			}
		}

		if doesTwo {
			containsTwoLetters++
		}

		if doesThree {
			containsThreeLetters++
		}
	}

	checkSum := containsTwoLetters * containsThreeLetters

	fmt.Println(checkSum)
}
