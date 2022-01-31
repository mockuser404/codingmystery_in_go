package main

import (
	"fmt"
	"math"
	"strings"
)

const particle runeT = runeTcount

func calibrateParticles() {
	downloadPuzzleInput("particle-grid.txt")
	grid, err := readPuzzleInput("particle-grid.txt")
	panicError(err)
	runeMap := map[runeT]rune{
		particle: '•',
		wall:     '█',
		player:   'C',
	}
	playerPos := pos{}
	numlines := numLines(grid)
	energyGrid := make([][]rune, numlines)
	foundPlayer := false
	for i, row := range strings.Split(grid, "\n") {
		if i < numlines {
			energyGrid[i] = []rune(row)
			if !foundPlayer {
				// TODO this can be removed and grid.w/2, grid.h/2 can be used
				// because in the question C is said to be the center.
				for j, r := range energyGrid[i] {
					if r == runeMap[player] {
						playerPos = pos{j, i}
						foundPlayer = true
					}
				}
			}
		}
	}
	d := float64(0)
	for y, row := range energyGrid {
		for x, r := range row {
			switch r {
			case runeMap[particle]:
				d += math.Abs(float64(playerPos.x-x)) + math.Abs(float64(playerPos.y-y))
			}
		}
	}
	fmt.Println(d, playerPos)
}
