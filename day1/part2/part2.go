package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	visited := make(map[string]int)
	direction := "N"
	locX := 0
	locY := 0

	// input := "R2, L3"
	// input := "R2, R2, R2"
	// input := "R5, L5, R5, R3"
	//input := "R8, R4, R4, R8"
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

		distance, _ := strconv.Atoi(m[1:len(m)])

		switch direction {
		case "N":
			for i := 1; i <= distance; i++ {
				locY += 1
				key := fmt.Sprintf("%d,%d", locX, locY)
				_, v := visited[key]
				if v {
					fmt.Println("visited key " + key)
					fmt.Println(abs(locX) + abs(locY))
					return
				} else {
					visited[key] = 1
				}
			}
		case "S":
			//locY -= distance
			for i := 1; i <= distance; i++ {
				locY -= 1
				key := fmt.Sprintf("%d,%d", locX, locY)
				_, v := visited[key]
				if v {
					fmt.Println("visited key " + key)
					fmt.Println(abs(locX) + abs(locY))
					return
				} else {
					visited[key] = 1
				}
			}
		case "E":
			//locX += distance
			for i := 1; i <= distance; i++ {
				locX += 1
				key := fmt.Sprintf("%d,%d", locX, locY)
				_, v := visited[key]
				if v {
					fmt.Println("visited key " + key)
					fmt.Println(abs(locX) + abs(locY))
					return
				} else {
					visited[key] = 1
				}
			}
		case "W":
			// locX -= distance
			for i := 1; i <= distance; i++ {
				locX -= 1
				key := fmt.Sprintf("%d,%d", locX, locY)
				_, v := visited[key]
				if v {
					fmt.Println("visited key " + key)
					fmt.Println(abs(locX) + abs(locY))
					return
				} else {
					visited[key] = 1
				}
			}
		}
	}

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
