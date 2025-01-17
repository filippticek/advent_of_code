package main

import (
	"fmt"
	"strconv"
	"strings"
)

var test1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var test2 = `Game 1: 10 green, 5 blue; 1 red, 9 green, 10 blue; 5 blue, 6 green, 2 red; 7 green, 9 blue, 1 red; 2 red, 10 blue, 10 green; 7 blue, 1 red
Game 2: 7 green, 5 red, 3 blue; 4 blue, 7 green, 8 red; 9 blue, 4 green; 6 green, 3 red, 4 blue
Game 3: 2 green, 4 blue, 13 red; 15 blue, 9 red, 3 green; 3 red, 18 blue, 3 green; 6 red, 4 green, 2 blue; 6 blue, 13 red
Game 4: 9 red, 1 green, 13 blue; 3 red; 2 blue, 6 red, 1 green; 12 blue, 2 red
Game 5: 1 red, 8 green; 2 red, 8 green, 8 blue; 1 red, 11 green; 5 blue, 11 green; 11 blue, 2 green; 10 blue, 2 red, 1 green
Game 6: 1 red, 12 blue; 20 blue, 3 green, 2 red; 4 red, 15 blue
Game 7: 13 red, 9 green, 1 blue; 8 green, 2 red, 6 blue; 4 green, 5 blue; 7 red, 3 green, 7 blue; 19 red, 5 blue, 1 green
Game 8: 11 red, 14 green, 4 blue; 2 blue, 5 green, 16 red; 18 blue, 11 red, 2 green; 2 blue, 15 red; 13 green, 8 blue
Game 9: 7 green, 5 blue, 11 red; 10 red, 7 green, 4 blue; 1 red; 6 green, 2 blue, 9 red; 8 green, 10 red, 6 blue; 5 red, 5 green, 7 blue
Game 10: 4 blue, 2 green, 1 red; 5 green, 2 red, 1 blue; 3 green, 8 blue, 1 red; 2 blue, 6 green, 2 red; 1 red, 4 green, 2 blue
Game 11: 3 red, 4 blue; 8 blue, 7 green, 2 red; 7 blue, 1 red, 6 green; 13 blue, 4 green
Game 12: 2 red, 3 blue, 4 green; 2 blue, 9 red, 8 green; 10 red, 1 blue; 1 green, 7 red, 3 blue; 7 red, 2 blue, 9 green
Game 13: 12 red, 6 green, 2 blue; 15 green, 2 red, 4 blue; 7 green, 1 red, 3 blue
Game 14: 9 green, 4 red; 6 blue, 1 red, 7 green; 3 blue, 5 green
Game 15: 7 red, 3 green, 2 blue; 3 blue, 4 green; 4 blue, 4 green, 9 red
Game 16: 12 blue, 11 green, 4 red; 8 blue, 9 red, 10 green; 9 green, 11 blue, 13 red; 10 red, 5 blue, 6 green; 2 red; 2 blue, 5 green, 5 red
Game 17: 3 red, 2 green, 2 blue; 1 blue, 3 red, 1 green; 10 green
Game 18: 3 green, 1 blue, 4 red; 12 red, 5 green; 3 red, 3 green, 3 blue; 13 red, 2 blue
Game 19: 13 blue, 8 green, 6 red; 10 red, 12 blue; 8 green, 13 red, 9 blue; 13 green, 3 red, 5 blue; 5 green, 1 blue, 2 red
Game 20: 19 red, 13 blue, 4 green; 1 red, 4 green, 8 blue; 14 red, 6 blue, 7 green; 11 red, 13 blue, 8 green
Game 21: 3 green, 13 red, 7 blue; 1 blue, 1 green, 1 red; 3 blue, 15 red, 5 green; 3 blue, 15 red, 2 green; 6 green, 9 red, 14 blue
Game 22: 2 red, 6 green, 4 blue; 6 green, 2 red; 1 blue, 4 red, 3 green; 11 green, 7 blue, 1 red; 4 red, 8 green, 3 blue
Game 23: 14 blue; 3 green, 2 red, 3 blue; 5 blue, 1 red
Game 24: 12 red; 5 blue, 16 red; 2 blue, 1 green, 16 red; 1 green, 11 red; 2 blue, 8 red, 1 green
Game 25: 4 red, 11 blue, 1 green; 7 red, 9 blue; 6 blue, 10 green
Game 26: 3 green, 13 red; 7 blue, 13 red, 5 green; 5 blue, 8 green, 11 red; 7 blue, 18 green, 6 red
Game 27: 6 green, 6 red, 5 blue; 2 blue, 4 green, 11 red; 15 red, 6 green; 4 green, 12 red, 2 blue; 3 blue, 5 red
Game 28: 16 blue, 6 red, 1 green; 7 red, 4 green, 10 blue; 1 red, 4 green
Game 29: 5 blue, 4 red; 6 blue, 3 red, 4 green; 2 green, 4 red, 5 blue; 1 green, 7 blue, 4 red; 3 green, 2 blue, 4 red
Game 30: 2 green; 14 green, 1 blue, 2 red; 5 red, 14 green
Game 31: 9 blue, 6 red, 7 green; 20 red, 1 green, 15 blue; 6 blue, 7 green, 17 red; 2 blue, 3 green, 6 red; 1 red, 3 blue, 2 green; 5 green, 18 red, 6 blue
Game 32: 7 green, 9 blue, 8 red; 8 red, 13 green, 19 blue; 2 red, 9 blue, 3 green; 9 green, 6 blue, 6 red
Game 33: 6 blue, 12 red; 13 blue, 3 green, 15 red; 5 red, 10 blue, 4 green; 11 blue, 6 red
Game 34: 5 green, 16 blue, 6 red; 10 green, 1 blue, 4 red; 2 red, 7 blue, 6 green; 12 green, 4 blue, 4 red
Game 35: 11 green, 3 blue; 1 red, 6 blue, 10 green; 11 green, 3 blue; 1 red, 2 blue; 11 green, 3 blue, 1 red; 3 blue, 2 green
Game 36: 10 green, 6 red, 4 blue; 3 green, 3 blue, 5 red; 6 red, 5 blue, 10 green
Game 37: 8 red, 7 blue, 5 green; 8 blue, 5 green, 14 red; 8 red, 2 blue
Game 38: 4 green, 1 red, 4 blue; 8 green, 11 blue; 7 red, 5 blue
Game 39: 2 blue, 4 red, 4 green; 8 green, 8 red, 1 blue; 3 red
Game 40: 4 green, 3 red, 14 blue; 4 blue, 13 green, 3 red; 12 green, 2 red, 2 blue; 8 green, 1 red, 11 blue; 4 green, 1 red, 1 blue
Game 41: 4 red, 1 green, 2 blue; 4 red; 4 red
Game 42: 12 red, 8 blue, 1 green; 8 green, 6 red, 5 blue; 12 green, 3 red, 13 blue; 1 red, 2 green, 8 blue; 3 green, 5 red, 6 blue
Game 43: 12 green, 3 blue; 13 green, 7 red, 5 blue; 10 green, 4 red, 4 blue
Game 44: 6 green, 2 red, 4 blue; 10 green, 6 red; 5 blue, 15 red, 13 green; 1 blue, 6 red, 3 green; 9 red, 5 green, 3 blue; 6 green, 4 blue, 5 red
Game 45: 10 blue, 14 green; 2 green, 2 red, 12 blue; 7 green, 1 red; 8 blue, 6 green, 1 red
Game 46: 8 red, 10 green, 15 blue; 9 green, 3 red, 17 blue; 2 blue, 10 red, 5 green; 11 blue, 3 green, 9 red; 5 red, 11 blue, 1 green; 7 green, 5 red, 16 blue
Game 47: 10 blue, 1 green, 1 red; 3 red, 8 blue, 7 green; 8 red, 9 blue; 2 green, 8 red, 1 blue
Game 48: 5 blue, 2 green; 2 red, 7 green, 2 blue; 1 blue, 3 green, 1 red
Game 49: 2 green, 6 red, 5 blue; 6 green; 4 blue, 17 red, 5 green
Game 50: 5 blue, 10 green; 6 blue, 4 red, 9 green; 7 red, 4 blue; 7 red, 3 blue, 14 green; 5 blue, 10 green, 9 red; 13 green, 1 blue, 9 red
Game 51: 1 blue, 15 green; 6 green, 2 blue; 5 blue, 1 red, 12 green
Game 52: 3 red, 15 green; 7 blue, 1 red, 14 green; 8 green, 1 red, 12 blue; 1 red, 9 green, 7 blue
Game 53: 2 green, 4 red; 1 red, 1 blue; 3 blue, 1 green; 2 red, 2 blue, 2 green
Game 54: 7 blue, 13 red, 7 green; 1 red, 2 green; 11 red, 10 green, 5 blue; 10 red, 8 green, 5 blue; 8 green, 12 blue, 12 red
Game 55: 18 red, 3 green, 5 blue; 5 green, 3 blue, 7 red; 3 blue, 3 green, 4 red
Game 56: 14 red, 17 green, 2 blue; 5 green, 13 red, 1 blue; 11 red, 20 green
Game 57: 3 red, 6 green, 2 blue; 3 red, 2 green; 2 green, 5 red; 1 blue, 1 green, 2 red
Game 58: 7 blue, 5 green, 9 red; 10 red, 5 green, 9 blue; 2 blue, 3 red, 8 green; 8 blue, 9 red; 7 red, 3 blue, 7 green; 2 green, 7 red, 1 blue
Game 59: 4 green, 3 blue; 10 red, 4 green, 4 blue; 2 green, 14 red, 12 blue; 1 blue, 1 green, 13 red; 10 red, 3 green, 3 blue; 2 green
Game 60: 9 red, 13 blue; 2 green, 5 red, 9 blue; 3 green, 10 blue
Game 61: 2 red, 8 green, 4 blue; 3 green, 2 red; 10 red, 9 green, 12 blue; 11 green, 17 blue, 3 red; 7 green, 1 red, 14 blue
Game 62: 1 green, 5 red, 13 blue; 5 blue, 1 green, 8 red; 2 green, 8 blue, 3 red; 1 green, 8 red
Game 63: 8 green, 15 red, 2 blue; 4 blue, 3 red, 12 green; 4 green, 1 blue, 17 red; 9 green, 18 red, 4 blue
Game 64: 7 blue, 17 red, 17 green; 3 blue, 4 green, 3 red; 4 red, 19 green, 1 blue; 11 blue, 14 red; 4 blue, 19 green, 7 red; 1 red, 10 green, 11 blue
Game 65: 1 blue, 17 red, 5 green; 17 red, 3 blue, 2 green; 10 blue, 9 green
Game 66: 5 blue, 6 red; 8 red, 2 blue, 1 green; 2 green, 3 blue; 8 blue, 10 red; 1 green, 2 red, 5 blue; 1 red, 3 blue
Game 67: 12 green, 16 blue, 12 red; 15 red, 1 blue, 3 green; 10 red, 3 green, 10 blue; 2 blue, 6 green, 6 red; 9 red, 8 blue, 7 green
Game 68: 10 red, 7 blue; 12 blue, 9 red; 12 blue, 9 red, 2 green
Game 69: 14 blue, 3 red, 3 green; 7 green, 7 red, 2 blue; 8 blue, 4 green, 8 red; 6 blue, 14 red, 3 green
Game 70: 7 blue, 6 green, 2 red; 2 red, 4 blue, 4 green; 2 red, 5 blue, 3 green; 6 green, 2 blue; 5 blue, 2 red, 2 green
Game 71: 7 green, 15 blue, 3 red; 15 blue, 15 red, 2 green; 10 red, 9 blue; 6 green, 20 blue, 11 red; 12 blue, 3 green, 7 red; 1 red, 7 blue
Game 72: 2 green, 9 blue, 7 red; 5 green, 3 blue, 5 red; 10 blue, 8 red, 7 green
Game 73: 18 blue, 5 red, 1 green; 18 blue, 3 red, 9 green; 2 red, 4 blue, 9 green; 5 blue, 5 red; 2 blue, 10 green, 6 red
Game 74: 1 blue, 10 green, 5 red; 4 green, 12 blue, 6 red; 7 red, 13 green, 3 blue; 5 blue, 8 green, 4 red
Game 75: 4 red, 2 blue, 5 green; 2 blue, 7 red, 4 green; 2 blue, 4 green, 3 red; 12 green, 2 blue; 10 green, 1 blue, 2 red
Game 76: 8 green, 6 blue, 5 red; 1 red, 2 blue, 9 green; 7 red, 9 green; 5 green, 1 blue, 11 red
Game 77: 3 blue, 10 red, 9 green; 7 blue, 6 red, 4 green; 4 red, 1 green, 8 blue
Game 78: 2 blue, 1 red, 14 green; 11 green, 1 blue; 15 green, 1 red
Game 79: 3 green, 17 blue, 1 red; 3 red, 2 blue, 10 green; 13 blue, 11 green, 5 red; 16 blue, 2 green, 16 red; 11 green, 1 blue, 14 red
Game 80: 7 red, 10 blue, 5 green; 6 blue, 6 green, 8 red; 6 blue, 3 green, 5 red
Game 81: 1 blue, 14 red, 6 green; 1 red, 13 blue, 12 green; 2 green, 15 red, 15 blue
Game 82: 5 blue, 8 red, 6 green; 19 blue, 4 green; 9 green, 15 blue, 3 red
Game 83: 19 red, 15 green, 2 blue; 17 red, 4 green, 1 blue; 13 green, 18 red
Game 84: 9 green, 14 red; 11 green, 14 red, 1 blue; 1 blue, 2 red, 3 green; 13 green, 10 red; 1 green, 5 red
Game 85: 4 red, 2 green, 11 blue; 8 blue, 3 red; 4 red, 1 blue, 5 green; 2 red, 3 green; 1 green, 8 red, 12 blue
Game 86: 5 blue, 1 red; 8 blue; 2 red, 1 green, 12 blue; 12 blue, 2 red
Game 87: 3 red, 10 green, 3 blue; 13 blue, 6 red, 2 green; 1 green, 2 red, 10 blue
Game 88: 10 red, 3 green, 8 blue; 3 red, 18 blue, 2 green; 3 green, 15 blue; 15 green, 16 blue, 8 red
Game 89: 10 blue, 1 red; 4 green, 9 red, 13 blue; 10 red, 3 green, 12 blue; 2 green, 1 red, 16 blue; 10 blue, 1 red, 6 green
Game 90: 4 red, 2 blue, 15 green; 5 red, 1 blue, 12 green; 3 blue, 3 red, 7 green; 4 blue, 3 red, 7 green; 1 red, 2 green, 1 blue; 1 blue, 4 green, 3 red
Game 91: 16 red, 10 blue, 1 green; 13 green, 13 red, 19 blue; 11 blue, 12 green, 2 red
Game 92: 8 blue, 4 green, 5 red; 7 blue, 4 red; 2 green, 15 blue; 16 blue, 4 red; 1 red, 7 green, 16 blue; 11 blue, 1 red, 3 green
Game 93: 12 green, 2 blue, 2 red; 8 red, 16 green, 8 blue; 15 red, 4 blue, 7 green; 1 red, 4 blue, 15 green; 13 green, 5 red, 4 blue; 5 green, 8 blue, 12 red
Game 94: 13 green, 10 red; 11 red, 19 green, 1 blue; 1 blue, 10 red, 12 green; 18 green, 9 red, 1 blue; 8 green, 1 red
Game 95: 3 green, 4 blue; 2 red, 2 green, 2 blue; 7 red, 3 green
Game 96: 5 red, 7 green; 4 blue, 14 green, 10 red; 13 green; 13 green, 3 blue; 13 green, 1 red, 3 blue; 12 red, 1 green
Game 97: 2 green, 1 blue; 9 red; 4 blue, 8 red; 4 green, 1 red, 14 blue; 2 green, 9 blue; 1 red, 6 blue, 2 green
Game 98: 12 green, 9 blue, 13 red; 6 red, 7 blue; 2 blue, 2 green
Game 99: 9 red, 3 green, 10 blue; 10 red, 10 blue, 4 green; 2 green, 15 blue, 3 red; 12 blue, 4 red
Game 100: 15 blue, 6 red; 1 green, 2 red; 12 blue, 8 green, 1 red; 1 red, 7 blue`

//
// 12 red cubes, 13 green cubes, and 14 blue cubes

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

func main() {

	games := strings.Split(test2, "\n")
	sum := 0
	for _, game := range games {
		sum += processGame(game)
		fmt.Println(sum)
	}
	fmt.Println(sum)

}

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func parseGame(line string) *Game {
	game := &Game{}
	slices := strings.Split(line, ":")
	id, _ := strconv.ParseInt(strings.Split(slices[0], " ")[1], 10, 32)
	game.ID = int(id)

	sets := strings.Split(slices[1], ";")
	for _, setString := range sets {
		set := Set{}
		colours := strings.Split(setString, ",")
		for _, colour := range colours {
			cMap := strings.Split(colour, " ")
			num, _ := strconv.ParseInt(cMap[1], 10, 32)
			switch cMap[2] {
			case "red":
				set.Red = int(num)
			case "green":
				set.Green = int(num)
			case "blue":
				set.Blue = int(num)
			}
		}
		game.Sets = append(game.Sets, set)
	}

	return game
}

func processGame(line string) int {
	game := parseGame(line)

	minRed := 0
	minGreen := 0
	minBlue := 0
	for _, set := range game.Sets {
		if set.Red > minRed {
			minRed = set.Red
		}
		if set.Green > minGreen {
			minGreen = set.Green
		}
		if set.Blue > minBlue {
			minBlue = set.Blue
		}
	}
	return minRed * minGreen * minBlue
}
