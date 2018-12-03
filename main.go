package main

import (
	"os"
)

func main() {
	day := os.Args[1]

	switch day {
	case "1":
		data := loadFromTextFile("day1.txt")
		dayOne(data)
	case "2":
		data := loadFromTextFile("day2.txt")
		dayTwo(data)
	case "3":
		data := loadFromTextFile("day3.txt")
		dayThree(data)
	}
}
