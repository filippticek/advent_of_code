package main

import (
	"fmt"
	"strings"
)

type Block struct {
	Vertical   []int
	Horizontal []int
}

func main() {
	part2(test2)
}

func part2(input string) {
	blocks := parse(input)

	sum := 0
	for _, block := range blocks {
		reflection := findReflection2(block)
		fmt.Println(reflection)
		sum += reflection
	}
	fmt.Println(sum)

}

func part1(input string) {
	blocks := parse(input)

	sum := 0
	for _, block := range blocks {
		reflection := findReflection(block)
		fmt.Println(reflection)
		sum += reflection
	}
	fmt.Println(sum)

}
func findReflection2(block Block) int {
	reflection := findSmudgedReflection(block.Vertical)
	if reflection > 0 {
		return reflection
	}
	return findSmudgedReflection(block.Horizontal) * 100
}

func findSmudgedReflection(direction []int) int {
	for i := 0; i < len(direction)-1; i++ {
		smudged := false
		smudge := 0
		if direction[i] == direction[i+1] || differAtOneBit(direction[i], direction[i+1]) {
			if differAtOneBit(direction[i], direction[i+1]) {
				smudged = true
				smudge = i
			}
			isDirection := true
			for j := 1; i-j >= 0 && i+j+1 < len(direction); j++ {
				if !smudged && differAtOneBit(direction[i-j], direction[i+j+1]) {
					smudged = true
					smudge = i
					continue
				}
				if direction[i-j] != direction[i+j+1] {
					isDirection = false
					break
				}
			}
			if smudged && isDirection {
				fmt.Println(direction, smudge)
				return (i + 1)
			}
		}
	}
	return 0

}
func differAtOneBit(a, b int) bool {
	x := a ^ b
	return x != 0 && ((x & (x - 1)) == 0)
}

func findReflection(block Block) int {
	fmt.Println(block.Horizontal, block.Vertical)
	for i := 0; i < len(block.Horizontal)-1; i++ {
		if block.Horizontal[i] == block.Horizontal[i+1] {
			isHorizontal := true
			for j := 0; i-j >= 0 && i+j+1 < len(block.Horizontal); j++ {
				if block.Horizontal[i-j] != block.Horizontal[i+j+1] {
					isHorizontal = false
					break
				}
			}
			if isHorizontal {
				return (i + 1) * 100
			}
		}
	}

	for i := 0; i < len(block.Vertical)-1; i++ {
		if block.Vertical[i] == block.Vertical[i+1] {
			isVertical := true
			for j := 0; i-j >= 0 && i+j+1 < len(block.Vertical); j++ {
				if block.Vertical[i-j] != block.Vertical[i+j+1] {
					isVertical = false
					break
				}
			}
			if isVertical {
				return i + 1
			}
		}
	}

	return 0
}

func parse(input string) []Block {
	blocks := make([]Block, 0)
	block := make([][]string, 0)
	i := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			blocks = append(blocks, parseBlock(block))
			i = 0
			block = make([][]string, 0)
			continue
		}
		block = append(block, make([]string, 0))
		block[i] = append(block[i], strings.Split(line, "")...)
		i++
	}
	return append(blocks, parseBlock(block))
}

func parseBlock(block [][]string) Block {
	blockStruct := &Block{Vertical: make([]int, len(block[0])), Horizontal: make([]int, len(block))}

	for j := 0; j < len(block[0]); j++ {
		for i := 0; i < len(block); i++ {
			if block[i][j] == "#" {
				blockStruct.Vertical[j] += (1 << i)
			}
		}
	}

	for i := 0; i < len(block); i++ {
		for j := 0; j < len(block[0]); j++ {
			if block[i][j] == "#" {
				blockStruct.Horizontal[i] += (1 << j)
			}
		}
	}

	return *blockStruct
}

var test1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var test2 = `#####..
.#..###
#..#.#.
#..#...
.#..###
#####..
##...#.
#.###..
.##....
#...###
#...###
.##....
#.###..
##...#.
#####..
.#..###
#..#...

..#..#..##..#
##.#..##..##.
..#.###.####.
.#.#.#...##..
.##.#.#..#.##
#####..#.....
#####..#.....
.##.#.#..#.##
.#.#.#..###..
..#.###.####.
##.#..##..##.
..#..#..##..#
..#..#..##..#

.##.#..##.###
..##..####.##
##.##.###.###
.........#...
##.#...##....
...#..#...#..
####.#.####..
##..##.....##
#####..#..###

...#...##...#..
..#....##....#.
..##.#....#.##.
##.....##.....#
##.##.#..#.##.#
##.##......##.#
#####.####.####
..#..######....
..#.########.#.

....#..#..#..
####.######..
....######.#.
...##.##.#.#.
##......#.#..
..##.##..#..#
......###.#..
##...##.##.##
##...##.##.##
......###.#..
..##.##..#..#
##......#.#..
...##.##.#.#.
.....#####.#.
####.######..
....#..#..#..
..##.#...#..#

###....
####..#
...###.
...####
##.....
##..##.
..#####
###....
....##.
##.#..#
##.####
##..##.
..#....
..##..#
..#....

.#.#..#.#.#..
...###.#...##
###..####..##
..##.##..####
#...#.....#.#
.##..#...#.##
#.##.##.#....
..#..#..#....
..#..#..#....

.###.#...##
.###.#...##
.##.#####..
##.#....#.#
#.###..####
#.###..####
##......#.#

#.##..#.###.#...#
.#.#....#....#...
.#.#.##.###...#.#
.#....####.#...##
..##..#.#...#..#.
..##..#.#...#..#.
.#....####.#...##
.#.#.##.###...#.#
.#.#....#....#...
#.##..#.###.#...#
..###.##.#.##..#.
##.#...#...#.#..#
##.#...#...#.#..#
..###.##.#.##..#.
#.##..#..##.#...#

#..#....#.#####
#..#.#......#..
...#.#......#..
#..#....#.#####
.######.....#..
..#.#.#..##.#.#
#..##.##...####
..#..##.##....#
.#......#.##..#
.###.##.###...#
####.#.######..
#####.#..###.##
..##.#...####..
..##.#...####..
#####.#..###.##

.#..#.#.##.
##...#..#.#
#.##.#.#.##
########..#
#....##.##.
#.##.#...##
.####.#....
.#..#.#.#..
##..##.#.#.
......#..##
#.##.#.#..#
.#..#...#.#
.#..#...#.#
#.##.#.#..#
......#..##
##..##.#.#.
.#..#.#.#..

...#..#...###.##.
...#......###.##.
..#..#####.#....#
##.#..#.#####.#..
....#.##...###..#
##.#..#.####.###.
..#.#..#....#.###

#.##.##......##
..##..##....##.
.#..#.##.##.##.
######.######.#
#....##.#..#.##
......##.##.##.
##...##.####.##
#######.#..#.##
#.##.#........#
..##..#......#.
#....##########

#.#.#....##....#.
#.....########...
..##..##.##.##.##
.##.##.######.##.
..#.#.#.#..#.#.#.
..#.#.###..###.#.
##.#.....##.....#
..###...####...##
#..##..#....#..##
#..##..#....#..##
..###...####...##

####..###....####
......#..##.##..#
.##..##..#....##.
#..#.#...#.##....
......#.#.#.##..#
....#.#######....
.##..#####....##.
####.#....#######
#..##...#...#####
#..########..#..#
.##.#...#...#....
####.#.....##.#..
####..#####..####
#..##...##..#####
#..##.#.#...#....
#..#.#####.##.##.
.....#####.######

###....
..##..#
..#####
###.##.
###.##.
....#..
...#..#

...#.#.
...#.##
.####..
###....
...##..
...##..
###....
.####..
...#.##
...#.#.
.##.#..
.#.#.##
.#.....
#.###..
.##.#..
.#.....
#.#..##

#....##.#.#..####
#.##.#....##.....
#....#...#...####
#.##.#......##..#
.#..#....#.##....
.#..#........#..#
#....######.##..#
......##.########
............#####
......####.######
#######.#.#...##.
##..####...#.#..#
######.###..#####
.......###.##.##.
#..#.#.#.###.####
.####.###.#######
##..######.##.##.

##..#####..##
#....#.#....#
#....#..####.
#.##.#.#....#
#.##.###....#
.#..#.#..##..
#######.####.
##..###.#..#.
.#.##.#......
..##..###..##
.####.#.#..#.
########....#
######..#..#.
#######..##..
......#.#..#.
.####....##..
#....##..##..

..####.#...
.#...####..
######.#.##
.....#.#.#.
..##..#.###
..#.###.#..
.###.##....
...###...##
...###...##
.###.##....
..#.###.#..

##..#.#
......#
##.#.##
..##.##
##.##.#
...#..#
..#.##.
..#..#.
...#..#
##.##.#
..##.##

..##...#..#..
##...###.#...
###.##.#.#..#
##.#...##.#..
....##.###...
###.###.##...
......##.#.#.
......##.#.#.
###.###.##...
....##.#.#...
##.#...##.#..

##.##........##
####.#..##..#.#
####.#..##..#.#
##.###......###
.#..#..#..#..#.
.....#......#..
....#.#.##.#.#.

.#.#.##..#.
.#..###.#.#
##..###.#.#
.#.#.##..#.
.##....##..
###.##..##.
###.##..##.
.##....##..
.#.#.##..#.
##..###.#.#
.#..###.#.#
.#.#.##..#.
#...##.###.

##..###
#.##.##
..##...
..##...
##..#..
#.##.##
.#..#..

####..#..#.#..#
.##.#####..#...
#.##.####.#.###
.###..#......##
...#######.##..
#.#.####..#..##
#.#.####..#..##
...#######.##..
.###..#......##
#.##.####.#####
.##.#####..#...
####..#..#.#..#
#.##.#..#.#.###
#.##.#..#.#.###
####..#..#.#..#
.##.#####..#...
#.##.####.#####

.##.###..##
....###..##
...#...##..
.##.#......
#.#####..##
..#########
###..##..##
.#.##..##..
###.#..##..
#.##...##..
#.###..##..
..##.......
.##.#......
##...######
##.#.##..##
####.#....#
...##......

###.##..#####...#
#...###....###...
.###.#.#..###...#
#.....#...##.###.
#.###..#...####..
###.#.###.##.#...
###.#.###.##.#...
#.###..#...####..
#.....#...##.###.
.###.#.#..###...#
#...###.....##...
###.##..#####...#
###..#..#.##.###.
....####.#.##.###
....####.#.##.###

#.###.#
.##...#
......#
.#.#.#.
.#.#.#.
.....##
.##...#
#.###.#
#.#.###
#.####.
#.####.

#.#..#.#.##.#
..####..####.
.######.####.
#......##..##
..#..#..#..#.
#.#..#.#.##.#
###..####..##
..#.##..####.
##....###..##
.#.##.#.####.
#.#..#.#....#
.#....#..##..
##.##.#######
###..########
.#....#.####.
##....##.##.#
.######......

..###...#.#
....###.###
##.###.##..
###.#..###.
..#.#.#.###
##.#######.
..####.....
####.###...
###.####...
##.#.#.##.#
...#.##.###
###.#..#.#.
###.#..#.#.
...#.##.###
##.#.#.####

####.#.###...
####.#.###...
..#.#.####...
#...##...#..#
#...#..#####.
.#.###.......
.#.####....#.
.....#..#.#.#
.#.#...#...#.
..##.##....#.
.###.#..#..#.
.###.#..#..#.
..##.##......
.#.#...#...#.
.....#..#.#.#

####...#....#
######.#....#
#..#.####..##
#..#..#.####.
.....#.#....#
####..###..##
.....####..##
#..####..##..
.##.##..#..#.
.##.#...#..#.
####...#####.

..##...
#....##
#....#.
..##...
#....#.
##..###
#....#.

.####.#..####
..##..#...##.
.####...#.##.
#....###.####
.#..#...#.#..
#.##.#...####
.####...#....
#.##.##..#..#
.####.#..#..#
.......#.....
..##..#...##.

###....##.#..
#...#..###...
##.#.###.....
####...##.#..
##........#.#
.#...#...#.##
##.##..#...#.
####...######
###.#...####.
..##.#....#.#
####.##.#.#.#
..##.######.#
#.##.######.#
#.##.######.#
..##.######.#

#######.#.####.
#.#.#.#.######.
#..#.#.#.####.#
#...###.##..##.
#..####...##...
.#..####.####.#
.#..#.#..####..
#.##.#...#..#..
##...#...#..#..
##...#...#..#..
#.##.#...#..#..
.#..#.#..####..
.#..####.####.#

..#..#..#..#..#
.#.##.#......#.
........#..#...
.........##....
........####...
..####...#....#
#......#.##.#..
.#....#..##..#.
#.#..#.##..##.#
##.##.###..###.
.######......##
.##..##..##..##
.######.#..#.##

.#....#..#....#..
#.####.###....###
.#.##.#.#......#.
..#..#....####...
.#.##.#...#..#...
#......#.##..##.#
###..####.####.##
.#....#..######..
##....##.######.#
##########....###
...##......##....
..####....####...
#.#..#.#...##...#
#.#..#.##..##...#
##....##.######.#
..#..#..##.##.##.
#......##########

.#.##.#.###
#.####.#...
#.#..#.##..
#.#..#.#.##
#..###.####
...##....##
###..###...

#.##.#..###.##.
.####.#.#..#..#
......#..##..##
#....###.##.#..
..##...#.#...##
#.##.##....##..
#.##.##...#.##.
#.##.##...#.##.
#.##.##....###.
..##...#.#...##
#....###.##.#..

.######.#.#.#
.######.#...#
###..###..#..
#..##..#.#..#
########.##..
..#..#..#....
#......###..#
##.##.##.####
########.####

##..#...###
..#.#.#####
###.#.##.##
####.###...
...#.###.##
....##..#..
###..#..#..
##.##.#.###
##..##..###
##..#.##...
....##.#..#

...#..#
...##..
..#..##
..#..##
...##..
.#.#..#
##...#.
##.##..
#....#.
.##.#..
.##.#..
#....#.
##.##..
##...#.
.#.#..#

...#...#......#..
....#...######...
##.###.##....##.#
##...#..#.##.#..#
...###.#.####.#.#
#####............
..#.#....#..#.#..
...###.#.#..#.#.#
###..#.#......#.#
##...#..#....#..#
##.#..#.#....#.#.

#.#.###....
#.#.###..#.
...########
...########
#.#.###..#.
#.#.###....
..#.#.#...#
##.######..
##....#..##
...##..###.
##.#..#...#
..#...#...#
.#.#..##...

##...##...##.##.#
...#....#...####.
###.#..#.########
#..........######
.#.#....#.#..##..
...##..##...####.
#.##....##..####.
..#.####.#..####.
..##.##.##..#..#.

##.#...####..#.##
###.#...##..#..##
#.#...#....##...#
#.#######.#.##...
.##....#.#...#.##
.#....#..##....##
#...##.###.#..###
###.##.###.#.####
.####..#...######
.####..#...######
###.##.###.#.####

.##...#..##..#...
####.##########.#
#####.########.##
.....###.##.##...
#..#.....##.....#
.....##..##..##..
#..#.###.##.###.#
########....#####
....##..#..#..##.
.##.###......###.
....#..#....#..#.
####.....##.....#
....##.######.##.

.#...#..#.####.#.
.....#....#..#...
#.##..#..#....#..
.#...#.#........#
.....###.#.##.#.#
.#...##.########.
#.#..............
###...##..#..#..#
###.#.##..#..#..#

.#...##
#.#.#.#
.##.#..
#.##..#
.#..##.
.#.###.
#.###..
.###...
...####
..##..#
#..####
.#####.
.##.##.
.##.##.
.#####.

...#..##..#####
####......##...
.##.##.....##..
######.####.#..
..#####.##.....
#......###..###
.#..#.#..###.##
.#........#....
......#......##
......#...#..##
.#........#....
.#..#.#..###.##
#......###..###

###......
###.#.##.
...#....#
...#....#
###.#.##.
###......
###.####.
.#.#....#
..##....#

...######
.....#...
....#....
..#...###
.########
###..####
....###.#
.#..#.###
###...###
#.###.#..
###..#.##
#...###..
.##.###..
..#.#.###
..#..#...
#..##.###
#..##.###

.##....##....
#..##.#..#.##
####.######.#
#..#..####..#
#####.####.##
#..###....###
######....###
.....#...##..
.##..##..##..
#..###.##.###
....#......#.
#####.#..#.##
#..#...##...#

.###.##....
###.#.#....
#####.#####
...###.....
.#.....####
.###...####
##.###..#..
.##.####..#
###..#..##.
#..###..##.
#..##.##..#
##..##.####
###.#...##.
...#.......
...#.......
###.#...##.
##..##.####

#.###..##
###..####
...##....
##.##.###
...##...#
##.##.##.
#.####.##
#.####.##
#......#.
#......#.
#.####.##

..############.#.
####.#...####.###
...#..#.####.#..#
..##.##...#..##.#
..##....#.#.#.###
..##....#.#.#.###
..##.##...#..#..#
...#..#.####.#..#
####.#...####.###
..############.#.
####...####.#.#..
##..###.##...#.##
####.#...##......

...####
#..#...
...####
..##.##
..##.##
...####
#.##...
...####
##.#.##
...#.##
#......

#.#..####
##.#..#.#
.#...#.#.
.#...#.#.
##.#..#.#
#.#..####
#.#....#.
#.#....#.
#.#..####
##.#..#.#
.#..##.#.

#....#...#.....
#....#..####..#
.#..#.##.###.##
.#..#..#.##...#
..##...##.#.#.#
.#.##...#.#...#
..##....#.##...
######.##..#.#.
######.##..#.#.

####.##.######...
####....####...##
#####..######..#.
####....######...
.##..##..##..##..
............#.##.
.............##.#
#..#.##.#..#..###
#############..#.
#..#....#..#.###.
#..#....#..#...##
.####..####.##..#
.##.####.##.##..#
....#..#......###
#..#.##.#..#.#.#.

.#.#..###
#..####..
####.#...
####.#...
#..####..
##.#..###
.#.#.#..#
####..###
.##......
..#..###.
..#..##..
..#..##..
..#..###.
.##......
####..###

#.##.##.......#..
#.##.#....###.#..
##..##.....######
.####...###...###
.####.#.#...#.###
.......####..#.##
.####.##...##..#.
..##......#..####
.####.####.##.###
##..###....##..##
.####.##...######
#....##.#..###...
.#..#..#.###.....

..##.##..##..
..####..##.##
...#.#.#.#.#.
##.#.........
..####...#.#.
##.#.#...##..
##.....####..
....###...###
..#.######.#.
..#.######.#.
..#.###...###
##.....####..
##.#.#...##..

.###.##.###.#.#..
####....####.#..#
#.#..##..#.##.###
.....##.......##.
.#.#.##.#.#.###.#
##..#..#..##....#
#.##....##.##.###
#.#.#..#.#.###...
..............##.
###############.#
###############.#
..............##.
#.#.#..#.#.###...
#.##....##.##.###
##..#..#..##....#
.#.#.##.#.#.###.#
.....##.....#.##.

.#.##.#
.#.##.#
.###.##
.###.#.
#.#..#.
##.##..
##.....
...#...
#.#...#
.#.#..#
#..####
#..#...
######.
###....
##.....
######.
#..#...

.#...#.#.......#.
.#.#.##..#..#.###
..###.##.....###.
##..#.###.#######
###.##.....##.#..
..##.#......##.#.
###...#..#.#....#
###...#..#.#....#
..##.#......##.#.
###.##.....##.#..
##..#.###.###.###
..###.##.....###.
.#.#.##..#..#.###
.#...#.#.......#.
.#...#.#.......#.

.....####..##
....#.#......
######..####.
#...#.##....#
.....#.......
#####.###..##
#..##.#.####.

.##.##...#.#.#.
.##.##...#.#.#.
..#.###..##....
.#.##.##..###..
#.##.##...##..#
####.####.....#
####.####.....#
#.##.##...##..#
##.##.##..###..

##.##....
..#......
...#.####
###.#....
##....##.
...######
##.#.##.#
.....#..#
..#......

#..#..########..#
.##.##.#.##.#.##.
.##.#...#..#...#.
######........###
.....####..#.##..
#..#............#
#..##.##....##.##
......###..###...
.##...##.##.##...

..#.##.
##..#.#
##..#.#
..#.##.
#####..
#.#....
..#####
####.##
.#..##.
..#####
.#.#.#.
.#..#.#
..##.#.
..##.#.
.#....#

#....##...#####
...........##..
##.###.####..##
#....###..#..##
######..####...
.#..#.##.##.#..
##..##.#####.##
##..###.####.##
..##...#...####

..##..##..#.##.
...####........
.########......
#.##..##.#.####
.#......#.#....
##.#..#.##.#..#
...####.....##.
..#....#..#.##.
###.##.###..##.
.#.#..#.#......
####..######..#
###.##.####.##.
#....#...#.....
.#.####.#..####
#..####..#.....

...##..#.#...#.
...##..#.##..#.
.###...#######.
..#..#..###..#.
.###...#...##..
..#.#..##.#...#
#..##.##...#...
##...#..#######
##...#..#######

......#.#
####..#.#
####.##..
##.....##
..#......
....####.
....#.##.

#.#.#.#
.###.#.
....#.#
....#.#
.###.#.
#.#.#.#
#.#####
.#..###
#..#.#.
#.##.#.
#.##.#.
#..#.#.
.#..###
#.##.##
#.#.#.#
.###.#.
....#.#

###.#.#..#.#.####
#.#.###..###.#.##
####..#..#..#####
##...##...#...###
..#.########.#...
...##.#..#.##....
.#.#..####..#.#..
##.#.#....#.#.###
###.#......#.####

...#......#..
......##.....
##..######..#
###..#..##.##
...########..
###.######.##
####.####.###
...#......#..
##....##....#
...#..##..#..
..###....###.

#...#...#.#...###
...#...#.##..##..
.###..#...#..####
#.##.##...##.####
..#.##..##.##.#..
###...#..##.##.##
...##.#.##.####..
...##.#.##.#.##..
###...#..##.##.##
..#.##..##.##.#..
#.##.##...##.####
.###..#...#..####
...#...#.##..##..

....#....#..#####
....#....#..#####
##..####..#.#.###
###...#.....##.##
....######.##..##
##.##...##...###.
.#..#####...#...#
....#.#..####..#.
..##.###.#.#.#...

#.#.#...#..
.##..###..#
#.####...#.
#.##.#...#.
.##..###..#
#.#.#..###.
........#.#
.#.#.#..#..
.#.#.#..#..
........#.#
#.#.#..###.
.##..###..#
#.##.#...#.
#.####...#.
.##..###..#

..##.##.##.##.##.
####..##..##..###
..##..######..##.
..#.##.#..#.##.#.
##..###.##..##..#
##..##########..#
##.#....##....#.#
####.#.#..#.#.###
....####..####...
..##..#....#..##.
.......####......
....#.#....#.#...
####..........###
..###.##..##.###.
###....####....##
####....##....###
###..#.####.#..##

#....###.
#...##.##
.#...##..
..#.#.##.
#.#..#.##
#.#..#.##
..#.#.##.
.#...##..
#...##.##
#....##..
.##.....#
#.##.###.
#..#.#..#
.#....#.#
.#....#.#

#..####
#..#.#.
#...###
##.####
....#..
.###.##
.#..#..
#.##.##
##.#...
..#..##
..#..##
##.#...
#.##.##

.##.#...#
##.#.##..
...#.....
########.
####..#.#
###...#..
.#..####.
#..#..##.
#####..#.
#####..#.
#..#..#..
#..#..#..
#####..#.
#####..#.
#..#..##.

......#...#.#....
.##.#####..##.###
#..#.##.#.##.###.
.##....#..#.#..##
.##....#..#.#..##
#..#.#..#.##.###.
.##.#####..##.###
......#...#.#....
....#..#..#..##.#
.##.#...##.#####.
.##..##.....##.##
.##.#####.#..#.#.
....#...#...###.#
#..#...#..####...
.##...##.#..###.#

#.####.##.#...##.
#.####.##.#...#..
###....######.#..
.#..###..#.######
.##.#..##..###.#.
.##.#..##..###.#.
.#..###..#.######
###....######.#..
#.####.##.#...#..
#.####.##.#...##.
###....#.#..#....
.###.#..##..##...
.####..#....##.##
#.#####.#####.##.
#..#...#..##.#...

..####.####..
.#....#.#..##
.###..#.##...
##.#.##...##.
##.#.##...##.
.###..#.##...
.#....#.#..##
..####.####..
##.##.#.###.#
##.##.#.#####
..####.####..
.#....#.#..##
.###..#.##...

##.##.#.#.#
.#.####..#.
###..##...#
##.####..##
.#..##.#.#.
.#..##.#.#.
##.####.###
###..##...#
.#.####..#.
##.##.#.#.#
#.#.#.###..
#.#.#.###..
##.##.#.#.#

###..##..##
.####..#.##
......#.#..
##.###.#.##
.#.#.#.#.##
...#...####
...#..##.##
.#.#.##....
#.####...##
###.##.##..
######.##..
#.####...##
.#.#.##....

.#..#..#...
#######.#..
#.##.#..#..
........###
#....###...
##..##.#...
#########..
##..##.##..
######.....
.......####
##..###.#..
......#.###
#...####...
#.##.##.#..
########.##

.....#...#..#
..#....#.####
#.###.#..####
..#.#....####
##..#.#.#####
####.#.##....
..#.####..##.
..##.#..##..#
##.##.####..#
####.#.#.#..#
....#..###..#
###.#.####..#
###.....##..#
##...#.#.####
..####..##..#

##...##....##..
......#..##..##
#..###.#.#..##.
..##.#.##..#.##
..##.#.##..#.##
#..###.#.#..##.
......#..##..##
##...##....##..
###.#..#....#..
.....#.###.#.#.
.#..#.#.###.#..
....#.#####....
.#...###....##.
...###...#.....
...###...#.....
.#...###....##.
...##.#####....

###...#..
##..##.##
.##....##
.#.#.####
...##....
....#..##
...###...
#...###..
.#.#..###
.....#...
#######..
######...
.....#...

..#..#..#..
#.##.#..#.#
#..##.##.##
#.###.##.##
#.##.#..#.#
..#..#..#..
..#.#.##.#.
.##.##..##.
##...#..#..

.#.#.####.#.#.##.
####......#######
####......#######
.#.#.####.#.#.##.
....######....#..
.#..........#.##.
.#..##..##..#.#..
##.##.##.##.##..#
####..##..####...
#.##..##..##.####
#...#....#.#.###.

#..##..##
.###.#...
#.#......
#.#......
.###.#...
#..##..##
...#.####
##.##..##
#.##.....
######...
....##.#.
#.###..##
##..#.###

##.####......
#####...##..#
.#..#.##...##
.#.#.##.#.#..
..##.#..#...#
#.########.#.
#.#######..#.
#.#######..#.
#.########.#.`
