package main

import (
	"fmt"
	"strings"
)

type runeT int

const (
	wall runeT = iota
	door
	player
)

type pos struct{ x, y int }

func navigateTunnels() {
	downloadPuzzleInput("map-of-the-tunnels.txt")
	downloadPuzzleInput("list-of-instructions.txt")

	tunnels, err := readPuzzleInput("map-of-the-tunnels.txt")
	panicError(err)
	instructions, err := readPuzzleInput("list-of-instructions.txt")
	panicError(err)

	startPos := pos{3, 21}

	numlines := numLines(tunnels)
	labrynthRunes := make([][]rune, numlines)
	for i, row := range strings.Split(tunnels, "\n") {
		if i < numlines {
			labrynthRunes[i] = []rune(row)
		}
	}

	runeMap := map[runeT]rune{
		wall:   '█',
		door:   '▒',
		player: 'I',
	}

	directions := []rune{}
	for _, r := range instructions {
		if r == ',' || r == ' ' || r == '\n' {
			continue
		}
		directions = append(directions, r)
	}
	dirmap := map[rune]pos{
		'E': {1, 0},
		'W': {-1, 0},
		'N': {0, -1},
		'S': {0, 1},
	}
	playerPos := startPos
	for _, dirRune := range directions {
		dir := dirmap[dirRune]
		next := pos{playerPos.x + dir.x, playerPos.y + dir.y}
		if labrynthRunes[next.y][next.x] == runeMap[wall] || labrynthRunes[next.y][next.x] == runeMap[door] {
			continue
		}
		playerPos = next
	}
	labrynthRunes[startPos.y][startPos.x] = 'S'
	labrynthRunes[playerPos.y][playerPos.x] = 'E'
	fmt.Println(playerPos)
}
