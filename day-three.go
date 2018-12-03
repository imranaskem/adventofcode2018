package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Claim represents a claim of cloth
type Claim struct {
	ID     int
	StartX int
	StartY int
	Width  int
	Height int
}

// Inch represents an inch of cloth
type Inch struct {
	Ids []int
}

func dayThree(data []string) {
	var claims []Claim

	for _, line := range data {
		nextClaim := NewClaim(line)
		claims = append(claims, nextClaim)
	}

	cloth := make([][]Inch, 1100)

	for i := range cloth {
		cloth[i] = make([]Inch, 1100)
	}

	for _, claim := range claims {

		for i := claim.StartX; i < (claim.StartX + claim.Width); i++ {

			for j := claim.StartY; j < (claim.StartY + claim.Height); j++ {
				if cloth[i][j].Ids == nil {
					cloth[i][j].Ids = make([]int, 0)
					cloth[i][j].Ids = append(cloth[i][j].Ids, claim.ID)
				} else {
					cloth[i][j].Ids = append(cloth[i][j].Ids, claim.ID)
				}
			}

		}
	}

	singleClaims := make(map[int]int)

	for _, x := range cloth {
		for _, y := range x {
			if len(y.Ids) == 1 {
				singleClaims[y.Ids[0]]++
			}
		}
	}

	for k, v := range singleClaims {
		if v == claims[k-1].GetArea() {
			fmt.Println(k)
		}
	}

	fmt.Println("End of Line")
}

// NewClaim acts as our constructor
func NewClaim(data string) Claim {
	re := regexp.MustCompile("[0-9]+")

	numbers := re.FindAllString(data, -1)

	id, _ := strconv.Atoi(numbers[0])
	startX, _ := strconv.Atoi(numbers[1])
	startY, _ := strconv.Atoi(numbers[2])
	width, _ := strconv.Atoi(numbers[3])
	height, _ := strconv.Atoi(numbers[4])

	claim := Claim{
		ID:     id,
		StartX: startX,
		StartY: startY,
		Width:  width,
		Height: height,
	}

	return claim
}

// GetArea retuns the area of a claim
func (c Claim) GetArea() int {
	return (c.Height * c.Width)
}
