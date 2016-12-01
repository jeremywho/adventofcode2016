package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	direction := "N"
	locX := 0
	locY := 0

	// input := "R2, L3"
	// input := "R2, R2, R2"
	// input := "R5, L5, R5, R3"
	input := `R4, R3, R5, L3, L5, R2, L2, R5, L2, R5, R5, R5, R1, R3, L2, L2, L1, R5, L3, R1, L2, R1, L3, L5, L1, R3, L4, R2, R4, L3, L1, R4, L4, R3, L5, L3, R188, R4, L1, R48, L5, R4, R71, R3, L2, R188, L3, R2, L3, R3, L5, L1, R1, L2, L4, L2, R5, L3, R3, R3, R4, L3, L4, R5, L4, L4, R3, R4, L4, R1, L3, L1, L1, R4, R1, L4, R1, L1, L3, R2, L2, R2, L1, R5, R3, R4, L5, R2, R5, L5, R1, R2, L1, L3, R3, R1, R3, L4, R4, L4, L1, R1, L2, L2, L4, R1, L3, R4, L2, R3, L1, L5, R4, R5, R2, R5, R1, R5, R1, R3, L3, L2, L2, L5, R2, L2, R5, R5, L2, R3, L5, R5, L2, R4, R2, L1, R3, L5, R3, R2, R5, L1, R3, L2, R2, R1`

	moves := strings.Split(input, ", ")

	for _, m := range moves {
		switch m[0] {
		case 'R':
			if direction == "N" {
				direction = "E"
			} else if direction == "E" {
				direction = "S"
			} else if direction == "S" {
				direction = "W"
			} else if direction == "W" {
				direction = "N"
			}
		case 'L':
			if direction == "N" {
				direction = "W"
			} else if direction == "W" {
				direction = "S"
			} else if direction == "S" {
				direction = "E"
			} else if direction == "E" {
				direction = "N"
			}
		}

		fmt.Println(direction)
		fmt.Println(m)
		fmt.Println(m[1:len(m)])
		distance, _ := strconv.Atoi(m[1:len(m)])
		fmt.Println(distance)

		switch direction {
		case "N":
			locY += distance
		case "S":
			locY -= distance
		case "E":
			locX += distance
		case "W":
			locX -= distance
		}
	}

	fmt.Println(moves)
	fmt.Println(locX)
	fmt.Println(locY)
	fmt.Println(abs(locX) + abs(locY))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
