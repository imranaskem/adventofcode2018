package main

import (
	"fmt"
	"strings"
)

func dayTwoPartOne(data []string) {
	linesOfText := data

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

func dayTwo(data []string) {
	linesOfText := data
	var commonLetters string
	idFound := false

	for _, originalLine := range linesOfText {

		for _, comparisonLine := range linesOfText {
			diffLetters := make(map[int]string)

			for i := range originalLine {
				if originalLine[i] != comparisonLine[i] {
					diffLetters[i] = string(originalLine[i])
				}

			}
			if len(diffLetters) == 1 {
				for k := range diffLetters {
					sb := strings.Builder{}
					sb.WriteString(originalLine[0:k])
					sb.WriteString(originalLine[k+1:])

					commonLetters = sb.String()
				}
				break
			}
		}

		if idFound {
			break
		}
	}

	fmt.Println(commonLetters)
}
