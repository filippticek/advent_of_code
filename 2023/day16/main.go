package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	RIGHT = 1
	DOWN  = 2
	LEFT  = 4
	UP    = 8
)

var (
	right = [2]int{1, 0}
	down  = [2]int{0, 1}
	left  = [2]int{-1, 0}
	up    = [2]int{0, -1}
)

func getDirection(direction [2]int) uint8 {
	return map[[2]int]uint8{
		right: RIGHT,
		down:  DOWN,
		left:  LEFT,
		up:    UP,
	}[direction]
}

func main() {

	part2("input")
}

func part2(file string) {
	field := parse(file)
	nums := []int{}

	for x := 0; x < len(field[0]); x++ {

		direction := down
		energized := make([][]uint8, len(field))
		for i := 0; i < len(field); i++ {
			energized[i] = make([]uint8, len(field))
		}

		moveLight(field, direction, energized, x, -1)

		sum := 0
		for _, row := range energized {
			for _, e := range row {
				if e != 0 {
					sum++
					continue
				}
			}
		}
		nums = append(nums, sum)
	}
	for x := 0; x < len(field[0]); x++ {

		direction := up
		energized := make([][]uint8, len(field))
		for i := 0; i < len(field); i++ {
			energized[i] = make([]uint8, len(field))
		}

		moveLight(field, direction, energized, x, len(field))

		sum := 0
		for _, row := range energized {
			for _, e := range row {
				if e != 0 {
					sum++
					continue
				}
			}
		}
		nums = append(nums, sum)
	}
	for y := 0; y < len(field); y++ {

		direction := right
		energized := make([][]uint8, len(field))
		for i := 0; i < len(field); i++ {
			energized[i] = make([]uint8, len(field))
		}

		moveLight(field, direction, energized, -1, y)

		sum := 0
		for _, row := range energized {
			for _, e := range row {
				if e != 0 {
					sum++
					continue
				}
			}
		}
		nums = append(nums, sum)
	}
	for y := 0; y < len(field); y++ {

		direction := left
		energized := make([][]uint8, len(field))
		for i := 0; i < len(field); i++ {
			energized[i] = make([]uint8, len(field))
		}

		moveLight(field, direction, energized, len(field), y)

		sum := 0
		for _, row := range energized {
			for _, e := range row {
				if e != 0 {
					sum++
					continue
				}
			}
		}
		nums = append(nums, sum)
	}

	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)

}

func part1(file string) {
	field := parse(file)

	direction := right
	energized := make([][]uint8, len(field))
	for i := 0; i < len(field); i++ {
		energized[i] = make([]uint8, len(field))
	}
	moveLight(field, direction, energized, -1, 0)

	sum := 0
	for _, row := range energized {
		for _, e := range row {
			if e != 0 {
				fmt.Printf("%2d", e)
				sum++
				continue
			}
			fmt.Print(" .")
		}
		fmt.Println()
	}
	fmt.Println(sum)

}

func moveLight(field [][]string, direction [2]int, energized [][]uint8, x, y int) {
	for {
		// for _, row := range energized {
		// 	for _, e := range row {
		// 		if e != 0 {
		// 			fmt.Printf("#")
		// 			continue
		// 		}
		// 		fmt.Print(".")
		// 	}
		// 	fmt.Println()
		// }

		x += direction[0]
		y += direction[1]
		if outOfBounds(x, y, len(field[0]), len(field)) {
			return
		}
		if wasDirectionHandled(direction, energized[y][x]) {
			return
		}

		fmt.Println("Y", y, "x", x, "CHAR", field[y][x], "ENERGIZED", energized[y][x])
		energized[y][x] += getDirection(direction)
		if energized[y][x] > 15 {
			fmt.Println("Y", y, "x", x, "CHAR", field[y][x], "ENERGIZED", energized[y][x])
			os.Exit(0)
		}
		switch field[y][x] {
		case ".":
		case "/":
			switch getDirection(direction) {
			case RIGHT:
				direction = up
			case DOWN:
				direction = left
			case LEFT:
				direction = down
			case UP:
				direction = right
			}
		case "\\":
			switch getDirection(direction) {
			case RIGHT:
				direction = down
			case DOWN:
				direction = right
			case LEFT:
				direction = up
			case UP:
				direction = left
			}
		case "-":
			switch getDirection(direction) {
			case RIGHT:
			case DOWN:
				moveLight(field, left, energized, x, y)
				direction = right
			case LEFT:
			case UP:
				moveLight(field, left, energized, x, y)
				direction = right
			}

		case "|":
			switch getDirection(direction) {
			case RIGHT:
				moveLight(field, up, energized, x, y)
				direction = down
			case DOWN:
			case LEFT:
				moveLight(field, up, energized, x, y)
				direction = down
			case UP:
			}
		}
	}
}

func wasDirectionHandled(direction [2]int, field uint8) bool {
	k := 0
	switch getDirection(direction) {
	case 1:
		k = 0
	case 2:
		k = 1
	case 4:
		k = 2
	case 8:
		k = 3
	}
	return ((field >> (k)) & 1) == 1
}

func outOfBounds(x, y, maxX, maxY int) bool {
	return x < 0 || y < 0 || x >= maxX || y >= maxY
}

func parse(file string) [][]string {
	content, err := os.ReadFile("./" + file)
	if err != nil {
		panic(err)
	}
	field := make([][]string, 0)
	for _, row := range strings.Split(string(content), "\n") {
		field = append(field, strings.Split(row, ""))
	}
	return field
}
