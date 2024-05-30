package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lense struct {
	Label       string
	FocalLength int
}

func main() {
	part2("input")
}

func part2(file string) {
	steps := parseFile(file)

	boxes := make(map[int][]*Lense, 0)

	for _, step := range steps {
		lense := parseLense(step)
		hash := hash(lense.Label)
		fmt.Println(lense)

		if lense.FocalLength > 0 {
			found := false
			for i := 0; i < len(boxes[hash]); i++ {
				if boxes[hash][i].Label == lense.Label {
					boxes[hash][i].FocalLength = lense.FocalLength
					found = true
				}
			}
			if !found {
				boxes[hash] = append(boxes[hash], lense)
			}
			continue
		}
		for i := 0; i < len(boxes[hash]); i++ {
			if boxes[hash][i].Label == lense.Label {
				boxes[hash] = append(boxes[hash][:i], boxes[hash][i+1:]...)
				continue
			}
		}
	}

	sum := 0
	for i, box := range boxes {
		for j, lense := range box {
			s := (i + 1) * (j + 1) * lense.FocalLength
			fmt.Println(s)
			sum += s
		}
	}
	fmt.Println(sum)
}

func parseLense(step string) *Lense {
	parts := strings.Split(step, "=")
	if len(parts) == 2 {
		return &Lense{Label: parts[0], FocalLength: int(getInt(parts[1]))}
	}

	parts = strings.Split(step, "-")
	return &Lense{Label: parts[0], FocalLength: 0}
}

func part1(file string) {
	steps := parseFile(file)

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println(sum)
}

func hash(seq string) int {

	sum := 0
	for _, c := range seq {
		sum += int(c)
		sum *= 17
		sum %= 256
	}

	return sum

}

func parseFile(file string) []string {
	content, err := os.ReadFile("./" + file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), ",")
}

func getInt(num string) int64 {
	n, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return n
	}
	return 0
}
