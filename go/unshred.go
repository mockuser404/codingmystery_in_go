package main

import (
	"fmt"
	"os"
	"sort"
)

/* func numRunes(s string) int {
	c := 0
	for range s {
		c += 1
	}
	return c
}

func lineLen(s string) int {
	c := 0
	for _, r := range s {
		c += 1
		if r == '\n' {
			return c
		}
	}
	return -1
}*/

func numLines(s string) int {
	lineC := 0
	for _, r := range s {
		if r == '\n' {
			lineC += 1
		}
	}
	return lineC
}

type strip struct {
	order int
	line  string
}

func shredRecovery() {
	downloadPuzzleInput("blank-sheet-of-paper.txt")
	downloadPuzzleInput("shredded-sheet-of-paper.txt")

	normal, err := readPuzzleInput("blank-sheet-of-paper.txt")
	panicError(err)
	tornStr, err := readPuzzleInput("shredded-sheet-of-paper.txt")
	panicError(err)

	fmt.Println(len(normal), len(tornStr))
	// fmt.Println(numRunes(normal) == numRunes(torn)) // assert
	// numRunes("█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ▀ ▒") => 20
	desiredLen := 20
	normalMap := map[string]int{}
	{
		stripIDbuf := make([]rune, desiredLen)
		lineNo := 0
		pos := 0
		for _, r := range normal {
			// i % lineNo should be == pos, not working with that logic
			if pos < desiredLen {
				stripIDbuf[pos] = r
			}
			if r == '\n' {
				key := string(stripIDbuf)
				normalMap[key] = lineNo
				lineNo += 1
				pos = 0
			} else {
				pos += 1
			}
		}
	}
	tornMap := map[string]int{}
	ss := make([]strip, numLines(tornStr))
	{
		stripIDbuf := make([]rune, desiredLen)
		linebuf := []rune{}
		lineNo := 0
		pos := 0
		for _, r := range tornStr {
			// i % lineNo should be == pos, not working with that logic
			if pos < desiredLen {
				stripIDbuf[pos] = r
			}
			linebuf = append(linebuf, r)
			if r == '\n' {
				key := string(stripIDbuf)
				tornMap[key] = lineNo
				// create a strip with order as defined in normalMap
				ss[lineNo] = strip{normalMap[key], string(linebuf)}
				lineNo += 1
				// reset
				linebuf = []rune{}
				pos = 0
			} else {
				pos += 1
			}
		}
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].order < ss[j].order
	})

	os.Mkdir("outputs", os.ModeDir)
	w, err := os.Create("outputs/" + "unshredded.txt")
	panicError(err)
	for _, strip := range ss {
		_, err = fmt.Fprintf(w, "%s", strip.line)
		panicError(err)
	}
	err = w.Close()
	panicError(err)
}
