package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SIZE = 256
)

type Edge struct {
	Direction string
	Moves     int
	RGB       string
}

func main() {
	part1("input")

}

func part1(file string) {
	edges := parse(file)

	field := make([][]bool, SIZE)
	for i := range field {
		field[i] = make([]bool, SIZE)
	}
	x, y := 0, 0
	for _, edge := range edges {
		x, y = move(field, edge, x, y)
		printField(field)
	}
	fillField(field)
	printField(field)
	fmt.Println(filledSum(field))
}

func filledSum(field [][]bool) int {
	sum := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] {
				sum++
			}
		}
	}
	return sum
}

func fillField(field [][]bool) {
	for d := 0; d < len(field)+len(field[0])-1; d++ {
		inside := false
		for x := 0; x < len(field[0]); x++ {
			y := d - x
			if y < 0 || y >= len(field) {
				continue
			}
			if field[y][x] {
				inside = !inside
				continue
			}
			if inside {
				field[y][x] = true
			}
		}
	}
}

func printField(field [][]bool) {
	fmt.Println()
	for _, row := range field {
		for _, c := range row {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func move(field [][]bool, edge Edge, x, y int) (int, int) {
	fmt.Println("X", x, "Y", y, edge.Direction, edge.Moves)

	switch edge.Direction {
	case "R":
		for i := 0; i <= edge.Moves; i++ {
			field[y][x+i] = true
		}
		return edge.Moves + x, y
	case "L":
		for i := edge.Moves; i >= 0; i-- {
			field[y][x-i] = true
		}
		return x - edge.Moves, y
	case "U":
		for i := edge.Moves; i >= 0; i-- {
			field[y-i][x] = true
		}
		return x, y - edge.Moves
	case "D":
		for i := 0; i <= edge.Moves; i++ {
			field[y+i][x] = true
		}
		return x, edge.Moves + y
	}
	return 0, 0
}

func parse(file string) []Edge {
	content, err := os.ReadFile("./" + file)
	if err != nil {
		panic(err)
	}
	edges := make([]Edge, 0)
	for _, row := range strings.Split(string(content), "\n") {
		parts := strings.Split(row, " ")
		edges = append(edges, Edge{
			Direction: parts[0],
			Moves:     int(getInt(parts[1])),
			RGB:       strings.Trim(parts[2], "()"),
		})
	}
	return edges
}

func getInt(num string) int64 {
	n, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return n
	}
	return 0
}
