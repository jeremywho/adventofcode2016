package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	x := 0
	y := 2

	mapper := [5][5]string{}

	mapper[0][0] = "X"
	mapper[0][1] = "X"
	mapper[0][2] = "1"
	mapper[0][3] = "X"
	mapper[0][4] = "X"

	mapper[1][0] = "X"
	mapper[1][1] = "2"
	mapper[1][2] = "3"
	mapper[1][3] = "4"
	mapper[1][4] = "X"

	mapper[2][0] = "5"
	mapper[2][1] = "6"
	mapper[2][2] = "7"
	mapper[2][3] = "8"
	mapper[2][4] = "9"

	mapper[3][0] = "X"
	mapper[3][1] = "A"
	mapper[3][2] = "B"
	mapper[3][3] = "C"
	mapper[3][4] = "X"

	mapper[4][0] = "X"
	mapper[4][1] = "X"
	mapper[4][2] = "D"
	mapper[4][3] = "X"
	mapper[4][4] = "X"

	// lines := `ULL
	//          RRDDD
	//          LURDL
	//          UUUUD`

	b, _ := ioutil.ReadFile("input.txt")
	lines := string(b[:])

	for _, line := range strings.Split(lines, "\n") {
		for _, c := range strings.TrimSpace(line) {
			switch c {
			case 'U':
				if y > 0 && mapper[y-1][x] != "X" {
					y--
				}
			case 'D':
				if y < 4 && mapper[y+1][x] != "X" {
					y++
				}
			case 'L':
				if x > 0 && mapper[y][x-1] != "X" {
					x--
				}
			case 'R':
				if x < 4 && mapper[y][x+1] != "X" {
					x++
				}
			}
		}

		fmt.Print(mapper[y][x])
	}

	fmt.Println("")
}
