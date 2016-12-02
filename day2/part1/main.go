package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	x := 0
	y := 0

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
				if y > 0 {
					y--
				}
			case 'D':
				if y < 2 {
					y++
				}
			case 'L':
				if x > 0 {
					x--
				}
			case 'R':
				if x < 2 {
					x++
				}
			}
		}

		num := (3 * y) + (x + 1)
		fmt.Print(num)
	}

	fmt.Println("")
}
