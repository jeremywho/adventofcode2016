package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	lines := string(b[:])

	total := 0
	count := 0
	for _, line := range strings.Split(lines, "\n") {
		total++
		if isValidTriangle(line) {
			count++
		}
	}

	fmt.Printf("Out of %d lines, %d where triangles\n", total, count)
}

func isValidTriangle(line string) bool {
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

	items := strings.Split(cleaned, " ")

	a, _ := strconv.Atoi(strings.TrimSpace(items[0]))
	b, _ := strconv.Atoi(strings.TrimSpace(items[1]))
	c, _ := strconv.Atoi(strings.TrimSpace(items[2]))

	return ((a + b) > c) && ((a + c) > b) && ((b + c) > a)
}
