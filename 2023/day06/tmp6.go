package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1(test1)
	part2(test2)
}

func part2(input string) {
	time := getInt(strings.ReplaceAll(strings.Split(strings.Split(input, "\n")[0], ":")[1], " ", ""))
	distance := getInt(strings.ReplaceAll(strings.Split(strings.Split(input, "\n")[1], ":")[1], " ", ""))
	fmt.Println(time, distance)

	numWins := int64(0)
	for j := int64(1); j < time; j++ {
		if j*(time-j) > distance {
			// fmt.Print(j, ",")
			numWins++
		}
	}
	fmt.Println("\nWins : ", numWins)
}

func part1(input string) {
	times, distances := parse1(input)
	fmt.Println(times, distances)
	wins := 0
	for i, time := range times {
		numWins := 0
		for j := 1; j < time; j++ {
			if j*(time-j) > distances[i] {
				fmt.Print(j, ",")
				numWins++
			}
		}
		fmt.Println("\nWins round: ", i, numWins)
		if numWins == 0 {
			continue
		}
		if wins == 0 {
			wins = numWins
			continue
		}
		wins *= numWins
	}

	fmt.Println(wins)
}

func parse1(input string) ([]int, []int) {
	s1 := getIntSlice(strings.Split(strings.Split(strings.Split(input, "\n")[0], ":")[1], " ")[1:])
	s2 := getIntSlice(strings.Split(strings.Split(strings.Split(input, "\n")[1], ":")[1], " ")[1:])
	return s1, s2
}
func getIntSlice(stringSlice []string) []int {
	intSlice := make([]int, 0)
	for _, num := range stringSlice {
		if n, err := strconv.ParseInt(num, 10, 32); err == nil {
			intSlice = append(intSlice, int(n))
		}
	}
	return intSlice
}

func getInt(num string) int64 {
	n, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return n
	}
	return 0
}

var test1 = `Time:      7  15   30
Distance:  9  40  200`

var test2 = `Time:        55     99     97     93
Distance:   401   1485   2274   1405`
