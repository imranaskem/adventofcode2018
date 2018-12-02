package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadFromTextFile(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
