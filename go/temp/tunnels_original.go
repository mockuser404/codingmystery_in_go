package temp

import (
	"fmt"
	"strings"
)

// remove these 4 to run the original one
func downloadPuzzleInput(s string)
func readPuzzleInput(s string) (string, error) {
	return "", nil
}
func panicError(err error)
func numLines(s string) int {
	return -1
}

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
	// fmt.Println(len(tunnels), len(instructions))

	startPos := pos{3, 21} // might change in the future so reading form map

	numlines := numLines(tunnels)
	labrynthRunes := make([][]rune, numlines)
	for i, row := range strings.Split(tunnels, "\n") {
		if i < numlines {
			labrynthRunes[i] = []rune(row)
		}
	}

	labrynthCopy := make([][]rune, numlines)
	data := make([]rune, numlines*len(labrynthRunes[0]))
	for i, row := range labrynthRunes {
		start := i * len(labrynthRunes[0])
		end := start + len(labrynthRunes[0])
		labrynthCopy[i] = data[start:end:end]
		copy(labrynthCopy[i], labrynthRunes[i])
		labrynthRunes[i] = []rune(row)
	}

	runeMap := map[runeT]rune{
		wall:   '█',
		door:   '▒',
		player: 'I',
	}

	fmt.Println(len(labrynthRunes), len(labrynthRunes[0]))
	x, y := 0, 0
	for y = 0; y < len(labrynthRunes); y++ {
		for x = 0; x < len(labrynthRunes[0]); x++ {
			// fmt.Print(string(labrynthRunes[y][x]))
			switch labrynthRunes[y][x] {
			// case runeMap[wall]:
			// case runeMap[door]:
			case runeMap[player]:
				startPos.x = x
				startPos.y = y
			default:
			}
		}
		// fmt.Println()
	}
	// fmt.Println(playerPos)

	directions := []rune{}
	for _, r := range instructions {
		if r == ',' || r == ' ' || r == '\n' {
			continue
		}
		directions = append(directions, r)
	}
	charmap := map[rune]rune{
		'E': '>',
		'W': '<',
		'N': '^',
		'S': 'v',
	}
	dirmap := map[rune]pos{
		'E': {1, 0},
		'W': {-1, 0},
		'N': {0, -1}, // N, S were reversed and wasted a ton of time
		'S': {0, 1},
	}
	playerPos := startPos
	for _, dirRune := range directions {
		/* for {
			next := pos{playerPos.x + dir.x, playerPos.y + dir.y}
			if labrynthRunes[next.y][next.x] == runeMap[wall] || labrynthRunes[next.y][next.x] == runeMap[door] {
				break
			}
			playerPos.x += dir.x
			playerPos.y += dir.y
			if next.y >= 0 && next.y < len(labrynthRunes) && next.x >= 0 && next.x < len(labrynthRunes[0]) {
				labrynthCopy[playerPos.y][playerPos.x] = '/'
			}
		} */
		dir := dirmap[dirRune]
		next := pos{playerPos.x + dir.x, playerPos.y + dir.y}
		if labrynthRunes[next.y][next.x] == runeMap[wall] || labrynthRunes[next.y][next.x] == runeMap[door] {
			continue
		}
		labrynthCopy[playerPos.y][playerPos.x] = charmap[dirRune]
		playerPos = next
		labrynthCopy[playerPos.y][playerPos.x] = 'I'
	}
	labrynthCopy[startPos.y][startPos.x] = 'S'
	// labrynthCopy[playerPos.y][playerPos.x] = 'E'
	for _, row := range labrynthCopy {
		fmt.Println(string(row))
	}
	fmt.Println(playerPos)
}
