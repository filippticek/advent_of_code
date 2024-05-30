package main

import (
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

const (
	MAX_ITER = 1000000000
	// MAX_ITER = 21
)

func main() {
	part2(test2)
}

func part2(input string) {
	matrix, rocks := parse2(input)

	allStates := make([]map[Point]bool, 0)
	printMatrix(matrix, rocks)
	for i := 0; i < MAX_ITER; i++ {
		rocks = moveNorth(matrix, rocks)
		// printMatrix(matrix, rocks)
		rocks = moveWest(matrix, rocks)
		// printMatrix(matrix, rocks)
		rocks = moveSouth(matrix, rocks)
		// printMatrix(matrix, rocks)
		rocks = moveEast(matrix, rocks)
		// printMatrix(matrix, rocks)

		if ok, oldState := isPattern(matrix, allStates, rocks, i); ok {
			lastState := (MAX_ITER-oldState)%(i-oldState) + oldState - 1
			fmt.Println("iteration", i, "old State", oldState, "lastState", lastState)
			for j := oldState; j < i; j++ {
				l := calculateLoad(matrix, allStates[j])
				fmt.Println("Load", l)
			}
			rocks = allStates[lastState]
			break
		}
		allStates = append(allStates, rocks)
	}

	load := calculateLoad(matrix, rocks)
	fmt.Println(load)
}

func isPattern(matrix [][]string, states []map[Point]bool, newRocks map[Point]bool, iter int) (bool, int) {
	if len(states) < 3 {
		return false, 0
	}
	for i, oldRocks := range states {
		found := true
		for rock := range oldRocks {
			if _, ok := newRocks[rock]; !ok {
				// fmt.Println("X", rock.X, "Y", rock.Y, "len", len(states))
				// fmt.Println("//////////////////////")
				found = false
				break
			}
		}
		if found {
			fmt.Println("OLD", i, "NEW", iter)
			// printMatrix(matrix, oldRocks)
			// printMatrix(matrix, newRocks)
			return true, i
		}
	}
	return false, 0
}

func printMatrix(matrix [][]string, rocks map[Point]bool) {
	for i, line := range matrix {
		for j, c := range line {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				fmt.Print("O")
				continue
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println("_____________________")
}

func calculateLoad(matrix [][]string, rocks map[Point]bool) int {
	loadSum := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				fmt.Print("O")
				loadSum += (len(matrix) - i)
				continue
			}
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("++++++++++++++++++++++")

	return loadSum
}

func moveNorth(matrix [][]string, rocks map[Point]bool) map[Point]bool {
	movedRocks := map[Point]bool{}
	load := make([]int, len(matrix[0]))
	for i := range load {
		load[i] = len(matrix)
	}
	loadSum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				loadSum += load[j]
				movedRocks[Point{X: j, Y: len(matrix) - load[j]}] = true
				load[j]--
				continue
			}
			switch matrix[i][j] {
			case ".":
				continue
			case "#":
				load[j] = len(matrix) - i - 1
			}
		}
	}

	return movedRocks
}
func moveWest(matrix [][]string, rocks map[Point]bool) map[Point]bool {
	movedRocks := map[Point]bool{}

	load := make([]int, len(matrix))
	for i := range load {
		load[i] = len(matrix[0])
	}
	loadSum := 0
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				loadSum += load[j]
				movedRocks[Point{X: len(matrix[0]) - load[i], Y: i}] = true
				load[i]--
			}
			switch matrix[i][j] {
			case ".":
				continue
			case "#":
				load[i] = len(matrix) - j - 1
			}
		}
	}
	return movedRocks
}
func moveSouth(matrix [][]string, rocks map[Point]bool) map[Point]bool {
	movedRocks := map[Point]bool{}
	load := make([]int, len(matrix[0]))
	for i := range load {
		load[i] = len(matrix) - 1
	}
	loadSum := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				loadSum += load[j]
				movedRocks[Point{X: j, Y: load[j]}] = true
				load[j]--
			}
			switch matrix[i][j] {
			case ".":
				continue
			case "#":
				load[j] = i - 1
			}
		}
	}

	return movedRocks
}
func moveEast(matrix [][]string, rocks map[Point]bool) map[Point]bool {
	movedRocks := map[Point]bool{}
	load := make([]int, len(matrix))
	for i := range load {
		load[i] = len(matrix[0]) - 1
	}
	loadSum := 0
	for j := len(matrix[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(matrix); i++ {
			if _, ok := rocks[Point{X: j, Y: i}]; ok {
				loadSum += load[j]
				movedRocks[Point{X: load[i], Y: i}] = true
				load[i]--
			}
			switch matrix[i][j] {
			case ".":
				continue
			case "#":
				load[i] = j - 1
			}
		}
	}

	return movedRocks
}

func parse2(input string) ([][]string, map[Point]bool) {
	matrix := make([][]string, 0)
	rocks := map[Point]bool{}
	for i, line := range strings.Split(input, "\n") {
		tmp := make([]string, 0)
		for j, c := range strings.Split(line, "") {
			if c == "O" {
				rocks[Point{X: j, Y: i}] = true
				c = "."
			}

			tmp = append(tmp, c)
		}
		matrix = append(matrix, tmp)
	}
	return matrix, rocks
}

func part1(input string) {
	matrix := parse(input)
	load := getLoad(matrix)
	fmt.Println(load)
}

func getLoad(matrix [][]string) int {
	load := make([]int, len(matrix[0]))
	for i := range load {
		load[i] = len(matrix)
	}
	loadSum := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			switch matrix[i][j] {
			case ".":
				continue
			case "#":
				load[j] = len(matrix) - i - 1
			case "O":
				loadSum += load[j]
				load[j]--
			}
		}
	}

	return loadSum
}

func parse(input string) [][]string {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		tmp := make([]string, 0)
		for _, c := range strings.Split(line, "") {
			tmp = append(tmp, c)
		}
		matrix = append(matrix, tmp)
	}
	return matrix
}

var test1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var test2 = `.#.OO.O.#OO........O....O..OO....OOO...#....#..........#.O.OO.....O#.#..#..#O##.O..#.#.O#.O.OO.....#
.O##.#..O..##..#..#.#..O#....O##..#.....##.......O...O#O..O.......#..#..O#O...#O..OO.#.....O.##.#.##
#.O...O.#.O....#.O..#OO...O...#.#.##...........#....#O.....##..#.O.#.............O...O..O.O...##...#
.#.#.....O..#...#...O.####OO....#.O...O...O...O.....OO#.O.##O.............O...#O.#OO#......O........
OO.O.....O....O..OO..O#..OO.OOO..OO...........O..##.O...O.#.O...##.O##....O.OO.#.O.....O..#.O...O##O
.#..##....#.O#...........#..OO...O......O.......#O..O..#...O......#.#....#.OO#..#O...O...#.....O..O#
##..OO.##.#........OOO##...OO.OOO#O...OOO.O.O...#...#.##..O##...#.#O#....#........#.O....#O.#O.O..O.
...O.O.O#...O.....O..OO...#O..#.#...O......#.......#.O...O.O#....O......#.OO.O..OOO.OO.O.O.###O...OO
...#O.....#O..##..#O.#...O.#...OO#O#.#..#.#....O..O.O.#...O.O...O.#..##O...O.#.OO.#O...OO.O.#O.OO#..
#.....#.#.O..#....O....##.O#......OO#....O#.#O.#..#O..#OO..#O..O...##O..#.....O##...O#.#..OO....O..#
.#OOO...#O...#...##...#O....#O..........O.#.O.O#..#..#OO.#............#....O#.#O....#....OO..#...O.O
#...OO...O...O..O.#.........O#O##..O........O.#..#O#O.##....#...#O.....OOO.#......#....O.O..#.OO.#..
.OO.....#..##...#.O.#O.#...##..#...#O.#...O#........O.O....#.##....O...#..#...#.O....OO..OO....O..#.
#..#.....OO....#O.#.O...O.OOO..O..O.#.O...OOOO..O.O...#..OOO...#......#.###.O.O....O.......O....O.#.
.#O.#....O....O#.OO..O.OOO...........O.##........O#O..#...O.#...O....O..O.##.#O#.#OO.O...O#...#..#O.
.OO.#O.#..O.O.OO...O.O....OOOO#.O.O.#.#.O.#.#....OO.#....#..#.....##...O....OO.OO...#.O.O...O.....OO
...###..O......#O.....#.....#...O.......O.#O.O..#..........OO..O.#..O...#..#.O.O....#..#.#...O.OOO..
........O..O.......O.......O#.O.O...#.#....#........#.#....OO....#.#........###.#...O...#OO.O#.#....
.OO#...........#......#.......#....O......#......##.O......OO.O#...O...O....O######O#..O..O.....OO..
.O.........OOO#O.#O...#.O....OOO##.......O.#..O..#..#O........O.OO..O..#...#O#..##..O#.###....OO....
.#.#.#O..#..#OO...O......O....OO..O......O......#..O..#...O.#..O...#....OOO....OO.OO#O....O.#.......
..OO..O.O..O#..#......OO.........#.O..O.O..#......#OO.#..O.O.#O#..O.#O...O.O..O.#......O...OOO.OO..#
O.O##....O.O....#.....#...OO..O..#.##..O..O###....#...O.#.O#O.OO.OO#.#...#.....#.#....O...#OOO..O#..
#.#.O.#O.O#.#...O..OOO.O.#.#......O..#O.#..##.##..#.........O...#OO...#..O.....#O...O..##O#O........
..O#O.##O.O.#.....#.OO..O.....OOOO#O....O..O.##O#..O..#..O#....##OO#..O...O..##..#O.##.........O#..O
#.#.#..##..O.O#O...##.#.O....#..#.O#..#....O.....#.#.....#.O.#O.....O.#.##...#..O.O.OOO..O...O......
....#....O..OO.#...O.O#.#.##..O....O..OO.#....O...O...O.....#......#..OO#OO.OO#..#.O#....#.#.OO..##O
O.##O.O.O....O..#......O..O.#O#O....#O....#.#.#.#O.........OO..O....O.#......O##OO......OOO..#O..#..
#.O.O.OO.O..#.....#O..#......O.......#....O.....O.....#...O.#.O......#...###...OO.#..O.OO#.O#..O#.OO
.....O.#.##OOOOOO#.#OO#O..#...#.........O..O.OO...OO....O..O.......O..O...#.#OO#.O.........O......O.
O......#..OO.O..O.......#..#...#..O.OO.#..##..#.O...#.O.#.O#O##.O.#.#..#..#....O.....#.#..O.....O#..
#.....#O...O.O......O......O.O.OO.O#O..#...O..#....#.....#O..O..#O...O..O.OO##.#.##O.#..#..#.#.....O
O#.O##..###.........O.....O.#.O#....#O.O..O.O#O#....OO..#O..O.....OO.O##..OO.#......#.O......###....
..O....#...OO.OO#.O#.O.OO#..O...##..OO..O#O#..OO..#.....OO...........O....#...#.......O#....O#......
....#.#O.O..OO......##.OO.#..##O.....#.......#.........O.O..OO.OO...#.#O.#O...OOOOO..O...OOOO.O#.#..
....#O##.OO#O..O.O.O#O.#O#..###....OO.....#OO..O.........O.O....OO.O..#..O#..OO...O#O.O.O.......O#..
O.O#.O.#.OO....OOO.O.O.O...OO.#.O.##..O..O.#..OO.#..O............#O.##.O#.OO.O.##..#....#..OO.O...#.
.#...O.#.#.#.O.#..##.O##.#..O..O...#.......O...OO.....#...#.OO#..O.OO....O....O.....#O.O#OO.##OO..O.
O#OO.OO..O..O.......O#.O....OO...O.........OO#.#O#O#.OO....#..#...O...OOO#...O.#O.#...........OO.O#.
..OO#..O...OO#O..#..#.#...#......O....O#..O.O....O#...##.....O..#O#..........O#.#......#.#.#OOO..O.O
.OO.......OO.#...O...O#.O...#..OO.##........O.........O..OOO.O..#......O#.OO.O....O.#..OOO....O..O.O
O#.O.#..O#..#O.#..O.O...##...O#.#O#...#O......O#.....#.###..#....O.....#O#....#....O#.O.O..#OO.O...O
.O...O...O...O.#.OO....#...O.O..........O...#...#........##..........##...#..O.O.O#O#.#......O..O...
.#.#..#O.#.....O.#....##........#....#....O.#O.O..............O.O....#..O.O.#O....O#O........#....O.
O....###O...O..O...OO.....#......O#.O..OO.#.#.O#....#O..#.O..OO..#..O.#.O..#....O...#OO.....O.OO.O..
#O.O.O.#O...#.OO...........#....O...#O.......#...##..##.##.O.O..........OOO..O..#..OO..O..OO..#.#.#.
#.O#..#..O.OOOO....OO.#O##.##..#..OO##O#.#.O....O.O#.##.O#...OO.O......#..#...O........O...#.##O..#.
......#...O......#......O.#..O.OOOO.O.O......#.#O...O#....O..#...........O.O..#...#..#O..O.O..##O..O
..O.#.....#...O.O..O....O.....#.....#.....O..O.O.O#O.#.....#.#OOOO...O.....O.#..#O..O.OOO....#OOO.O.
.OOO....O#.........O.......#..O......#..........#...##O...O...O...O.O.O##....O#O##.......##O...#..O.
#.##.O.O....O......##..#.....##..##..O#....O.#O#.O....#........#.#.#....OO.O.........OO..O......#O#.
#.O.......#O#.O....O.O.......O.O##O#....#O.O.#O#.O.O.OO..#.O#...O.##.###.O.O#.O..#...#O.....OO....##
..O......O...#..#.O#....##O......OO......O.O....O....O#...O.O.OOO.#O...O.#..O......OO...#.O#.....#.O
.OO...#O.O....O.........##......##.......#...O..##.....#...OO.....O...#.#....O..O...O#.......O..O.O.
.O#.....O....O..O...#.O.O#...#.#..#O#.O.....O#O.OO....#....O..OO.OO.#..#.#.O.O.....O..O......O..#..#
.O.#...O...O.OOO#O...#........O..#..#.##..#...#...#..O....##.O.O#...O.....O.##.O........O...#O#O#.O.
.#O....#.O#O##O......O........#...##.#..O.OO..O.....#.#.O#...O.O#.#...#.......#...#.#.O.....O.OOO...
.O....OO###...O#...O..#..O............#O.#.O.......OO...O...##......O.#..O.O#.#....O..O.#....O...##O
#..#O#.O...O....OO..O..........OO##.....O......O...##O.#..O..OO#...#.OO..#...#.O..O...O.O#.....#....
O.O#OO.....O#.O..##.OOO.##....O#...O..#.#...O..OO......O#..O..O..#...#O........##.O##....#O.....OO.O
......O..O..O....OOO..O...O.O#..#.....O..O..#....#.O##....O.#..##O...O..O..OO....#...O.#..O..OO...#.
..........O.#O#.#......#......#O..O.O.OOOOO.#...O..#.O.OO....#.O#.#O..#......#.#O.O.O...O..##..O#.O.
OO#...#...OOO..O.#.......#.O...##..O.O......#..O.......##.....O..O.....OOO.##..O..O.#...#O.......#.#
....##..#..O......OO.......O#O#O...O.O.#.#..#..##..##O....O.#.........#.O...#..O........O..#..#O.OO.
......O.O...OO##..O..##.O#O#..O##.O#..#...O##....#.......#....#OO.O.####O...###O...#.........###O..O
.#.O#.#O..#..#...#O..O..#..........#.......O..#.O...O.#..#O..O...O..#..O.............#.OO....O#..OO.
.....#.O.OO...#...O...O..........#......#..O.O....OO.#..#...O#O..#.....OO........O.#O.#.......OO#.OO
OO...#.O.......#O.O...O#.O....O.....OO.O.........##OO...OO...O.#..OO...O.#.O...O.O..O#OO.O#O..#..OO.
..O##.#..#...O..O.#....O.O##.O.#.O...O..#O.OO...O.#.....##O...O.O.......##O...OO..O#..O#.#......O.O.
.....O.....OO..O..O..O.#...O......O..OO..OO....OOO#.O..OO......#.O.....#O.O.O#....##OO..#O#.O..O#O.O
......O.O#...O..O.OO.O.OO.OOO...#O..O#...OO.#O.OO...O..#..O..O.#....O..OO....#......O.O.#O.##..O..O.
#.....#....O#O..O#O....O.......#.....O....#..#.OOOO#OOO#..#.OO....OO.O#.O......#..O.OO..O.#....O.O..
.O.O.#.#.O...#..#O.......##...OO..O...#......#.O#...O..O....#O.O.O.........#O.O#..O....##.O.O.......
....#..O.#.......#.#O..O..........#...#..O##..O#..O..OO.O###.#O#..O#O#O..#OO.#..O#.#.O..#....O.O....
...##...........O...#O.OOO..##O.O...#.OO.OO..O...O.....#.......#.....#........#....##.....O...#.#..O
.O......O.#...##O..#O..O......O#O....#.......O....#...O#....O.....#......#O..#.O..#.#.O.O#...O#.....
O..O.#........O.##.#..##....O.OO.......O.......O..#O#O......O...###.........OO...#.#..#........#.#..
O.........#..O.O...#O......#.O.#OO..O.##..O.#O.#..#..#.#.O.....#...#.....##....#.......O##...O#..O..
#OO..........O.O...#.OO....#...#.O#..OO#..#O.##.O....#O#..O....O..#...#.#....O.O...#.#O....OO.#.##..
.#...#O.#O#O###.OO.#..#....#OO.O....O.O#.O...........O#....O.#.O......#.O.#.#O.#..O.OO.....OO...OO.#
O#.........O....O.#....##.....O.....O.........#O.O#O...O.O#..O.#O....#.#...O....O#....##.......#....
.....OO#..O##..O.......O.#...#.OO.O.........#..O###...#.O..O#.O......O.....#.O.OO......#O...O#.....O
.O.O##.###OO..........O.......OOO..O.OO..##...#O..O...OO..#.O.....O.#..#.O.OO##.#.#...###....O...#..
.O..###...O...#.OOO#.O#.O#...O.#OO.OO...O..#.O......#.O.O.OO...#OO..#...#..O#O...O........##.OO.#...
.#..#O......#O......O#...#..#O...##O..#.O..O....O.#O..OO..O..#.#....O..O......O.O.#...##..O.O.#O....
.O.##.#...#O..#...O#.#.O..O#.......O....#.O.O.#.##.....#.......O.##..#.......O.....##.O.#.....#..#O.
...#.O.O.#O#.#...#...#OO.#O..O....O...O#....#.#.O..O..OO.#.OO..O...#O....O..#.......OO...O.....#.O..
O..O...O.......O.....##O#.O.#...OOO.O...........#..O..O.O.....O............#...........#...O.O..#.O.
..O.O.#O...O#.#.#O.OO.....##..O##O.O.O....#OOO...##.O.O.OO........O.#O........O.#.O.#.O.OO....O#.OO.
#.O.O##.....OO..#OO.O...O.#..#......O...OO....OO.#...O.##.O.OO.#.##..##O.O...#O#...##.....#..O#O.##.
..#OO..OO.##..#....O##.#O.OO##.O....#....O...O.........O....O#O.##.OOO.#..#......OO#..OO...#..#..#O.
.....#.#.O...O##....O.#O..O.#..#..O.#O.#O##..O.#..#O.#.O..O.......##......O#.##.O#..OO.....O#OOO.#..
O....O.O#.#...O.OO.........O...#...#...OO..##..OO..........#O.O.....#.#.OO.#..O#O.#...OO..O##.#.O...
..#....##..............O..#.#..OO.OOO..#OO.O#O.OO#..#...O.#...O..O...OO......#.O..#....O.OO.O....##.
O.#O##....#.OO....#.#...#.....O......O......#.#......#OO.....###...#...#..O...O...O.OO.#...#.......#
....O..#....#....O...#...O....OO.O##..#.O..OO.#.#.#......#O.OO......O##O.#O...#.#.....O....#O..O....
#OO.#..O......#.....#........#...#.....O...OOO.......OO##O....#......O.O...O##.O.OOO.#........O.#...
O.OO.......O..OO.O.OO.OO..OO#...##.#.....#...O.#.O....O...#.OO.##.#.....OO..O#......#OOOO....O....##
...#...#..#....#.....O..O#O..#.##O.O......O.#O......O..##...#O#O..##..O.#....O.#...O.OO...#O#....OOO
.........OOOO.O.O.O...#........#....OO.#OO#....#O........#.....#O.....##O.....#O..O.O...#...#....OO.`
