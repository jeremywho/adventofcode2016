package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(b[:]), "\n")

	total := 0
	count := 0
	for i := 0; i < len(lines); i = i + 3 {
		l1 := parse(lines[i])
		l2 := parse(lines[i+1])
		l3 := parse(lines[i+2])

		total += 3

		if isValidTriangle(l1[0], l2[0], l3[0]) {
			count++
		}

		if isValidTriangle(l1[1], l2[1], l3[1]) {
			count++
		}

		if isValidTriangle(l1[2], l2[2], l3[2]) {
			count++
		}
	}

	fmt.Printf("Out of %d lines, %d where triangles\n", total, count)
}

func parse(line string) []string {
	// do some weird stuff here to parse the numbers correctly...there's lots of spaces in the data
	cleaned := ""
	line = strings.TrimSpace(line)
	for i, c := range line {
		if c != ' ' {
			cleaned += string(c)
			continue
		}

		if c == ' ' && i < len(line) && line[i+1] == ' ' {
			continue
		} else {
			cleaned += string(c)
		}
	}

	return strings.Split(cleaned, " ")
}

func isValidTriangle(x, y, z string) bool {
	a, _ := strconv.Atoi(strings.TrimSpace(x))
	b, _ := strconv.Atoi(strings.TrimSpace(y))
	c, _ := strconv.Atoi(strings.TrimSpace(z))

	return ((a + b) > c) && ((a + c) > b) && ((b + c) > a)
}
