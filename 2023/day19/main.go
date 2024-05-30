package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StatePart struct {
	Part  PartRange
	State State
}
type PartRange struct {
	XMin int
	XMax int
	MMin int
	MMax int
	AMin int
	AMax int
	SMin int
	SMax int
}

type State struct {
	Name       string
	Conditions []Condition
}

type Condition struct {
	Category  string
	Less      bool
	Value     int
	NextState string
}

type Part struct {
	X int
	M int
	A int
	S int
}

func main() {
	part2("input")
}

func part2(file string) {
	states, _ := parse(file)
	states["A"] = State{Name: "A"}
	states["R"] = State{Name: "R"}

	statePart := StatePart{State: states["in"], Part: PartRange{
		XMin: 1,
		XMax: 4000,
		MMin: 1,
		MMax: 4000,
		AMin: 1,
		AMax: 4000,
		SMin: 1,
		SMax: 4000,
	}}
	stateParts := []StatePart{statePart}
	finishedParts := []PartRange{}
	for {
		stTemp := make([]StatePart, 0)
		fmt.Println("______________")
		for _, sp := range stateParts {
			fmt.Println(sp)
		}
		fmt.Println("______________")
		for _, statePart := range stateParts {
			sps := getPartState(states, statePart)

			for _, sp := range sps {
				fmt.Println("+++++++++++++++")
				fmt.Println(sp)
				if sp.State.Name == "A" {
					finishedParts = append(finishedParts, sp.Part)
				} else if sp.State.Name == "R" {

				} else {
					stTemp = append(stTemp, sp)
				}
			}
		}
		if len(stTemp) == 0 {
			break
		}
		stateParts = stTemp
	}

	sum := 0
	for _, fp := range finishedParts {
		fmt.Println(fp)
		sum += (fp.XMax - fp.XMin + 1) * (fp.SMax - fp.SMin + 1) * (fp.AMax - fp.AMin + 1) * (fp.MMax - fp.MMin + 1)

	}
	fmt.Println(sum)
}

func getPartState(states map[string]State, statePart StatePart) []StatePart {
	stateParts := make([]StatePart, 0)
	rng := &statePart.Part
	// fmt.Scanln()
	fmt.Println("((((((((((((((")
	fmt.Println(statePart)
	fmt.Println("))))))))))))))")
	for _, condition := range statePart.State.Conditions {
		nr := *rng
		nextRange := &nr
		switch condition.Category {
		case "x":
			if condition.Less {
				if rng.XMax < condition.Value {
					rng = nil
				} else if !(rng.XMin < condition.Value) {
					nextRange = nil
				} else {
					nextRange.XMax = condition.Value - 1
					rng.XMin = condition.Value
				}
			}
			if !condition.Less {
				if rng.XMin > condition.Value {
					rng = nil
				} else if !(rng.XMax > condition.Value) {
					nextRange = nil
				} else {
					nextRange.XMin = condition.Value + 1
					rng.XMax = condition.Value
				}
			}
		case "m":
			if condition.Less {
				if rng.MMax < condition.Value {
					rng = nil
				} else if !(rng.MMin < condition.Value) {
					nextRange = nil
				} else {
					nextRange.MMax = condition.Value - 1
					rng.MMin = condition.Value
				}
			}
			if !condition.Less {
				if rng.MMin > condition.Value {
					rng = nil
				} else if !(rng.MMax > condition.Value) {
					nextRange = nil
				} else {
					nextRange.MMin = condition.Value + 1
					rng.MMax = condition.Value
				}
			}
		case "a":
			if condition.Less {
				if rng.AMax < condition.Value {
					rng = nil
				} else if !(rng.AMin < condition.Value) {
					nextRange = nil
				} else {
					nextRange.AMax = condition.Value - 1
					rng.AMin = condition.Value
				}
			}
			if !condition.Less {
				if rng.AMin > condition.Value {
					rng = nil
				} else if !(rng.AMax > condition.Value) {
					nextRange = nil
				} else {
					nextRange.AMin = condition.Value + 1
					rng.AMax = condition.Value
				}
			}
		case "s":
			if condition.Less {
				if rng.SMax < condition.Value {
					rng = nil
				} else if !(rng.SMin < condition.Value) {
					nextRange = nil
				} else {
					nextRange.SMax = condition.Value - 1
					rng.SMin = condition.Value
				}
			}
			if !condition.Less {
				if rng.SMin > condition.Value {
					rng = nil
				} else if !(rng.SMax > condition.Value) {
					nextRange = nil
				} else {
					nextRange.SMin = condition.Value + 1
					rng.SMax = condition.Value
				}
			}
		default:
			if condition.NextState == "R" {
				return stateParts
			}
			if condition.NextState == "A" {
				return append(stateParts, StatePart{
					Part:  *nextRange,
					State: states[condition.NextState],
				})
			}
		}
		if nextRange == nil {
			continue
		}

		stateParts = append(stateParts, StatePart{
			Part:  *nextRange,
			State: states[condition.NextState],
		})
		if rng == nil {
			return stateParts
		}
	}
	return stateParts
}

func part1(file string) {
	states, parts := parse(file)
	for _, state := range states {
		fmt.Println(state.Name, state.Conditions)
	}
	for _, part := range parts {
		fmt.Println(part)
	}

	sum := 0
	for _, part := range parts {
		sum += handlePart(states, part)
	}
	fmt.Println(sum)

}

func handlePart(states map[string]State, part Part) int {
	state := states["in"]

	for {
		var nextState string
		for _, condition := range state.Conditions {
			switch condition.Category {
			case "x":
				if condition.Less && part.X < condition.Value {
					nextState = condition.NextState
				}
				if !condition.Less && part.X > condition.Value {
					nextState = condition.NextState
				}
			case "m":
				if condition.Less && part.M < condition.Value {
					nextState = condition.NextState
				}
				if !condition.Less && part.M > condition.Value {
					nextState = condition.NextState
				}
			case "a":
				if condition.Less && part.A < condition.Value {
					nextState = condition.NextState
				}
				if !condition.Less && part.A > condition.Value {
					nextState = condition.NextState
				}
			case "s":
				if condition.Less && part.S < condition.Value {
					nextState = condition.NextState
				}
				if !condition.Less && part.S > condition.Value {
					nextState = condition.NextState
				}
			default:
				nextState = condition.NextState
			}

			if nextState == "R" {
				return 0
			}
			if nextState == "A" {
				break
			}
			if nextState != "" {
				break
			}
		}
		if nextState == "A" {
			break
		}
		state = states[nextState]
	}

	return part.X + part.A + part.M + part.S

}

func parse(file string) (map[string]State, []Part) {
	content, err := os.ReadFile("./" + file)
	if err != nil {
		panic(err)
	}
	states := make(map[string]State, 0)
	parts := make([]Part, 0)

	lines := strings.Split(string(content), "\n")
	i := 0
	for ; lines[i] != ""; i++ {
		name := strings.Split(lines[i], "{")[0]
		conditionsStr := strings.Split(strings.TrimSuffix(strings.Split(lines[i], "{")[1], "}"), ",")
		conditions := make([]Condition, 0)
		for j := 0; j < len(conditionsStr)-1; j++ {
			cond := Condition{
				Category:  string(conditionsStr[j][0]),
				Less:      strings.Contains(conditionsStr[j], "<"),
				NextState: strings.Split(conditionsStr[j], ":")[1],
			}
			if cond.Less {
				cond.Value = int(getInt(strings.Split(strings.Split(conditionsStr[j], "<")[1], ":")[0]))
			} else {
				cond.Value = int(getInt(strings.Split(strings.Split(conditionsStr[j], ">")[1], ":")[0]))
			}
			conditions = append(conditions, cond)
		}
		conditions = append(conditions, Condition{Value: -1, NextState: conditionsStr[len(conditionsStr)-1]})
		states[name] = State{
			Name:       name,
			Conditions: conditions,
		}
	}
	i++
	for ; i < len(lines); i++ {
		line := strings.Trim(lines[i], "{}")
		cats := strings.Split(line, ",")
		part := Part{}
		for _, cat := range cats {
			if strings.Contains(cat, "x") {
				part.X = int(getInt(strings.Split(cat, "=")[1]))
			}
			if strings.Contains(cat, "m") {
				part.M = int(getInt(strings.Split(cat, "=")[1]))
			}
			if strings.Contains(cat, "a") {
				part.A = int(getInt(strings.Split(cat, "=")[1]))
			}
			if strings.Contains(cat, "s") {
				part.S = int(getInt(strings.Split(cat, "=")[1]))
			}
		}
		parts = append(parts, part)
	}

	return states, parts
}

func getInt(num string) int64 {
	n, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return n
	}
	return 0
}
