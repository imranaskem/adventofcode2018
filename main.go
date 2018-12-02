package main

import (
	"os"
)

func main() {
	day := os.Args[1]

	switch day {
	case "1":
		dayOne("day1.txt")
	case "2":
		dayTwo("day2.txt")
	}
}
