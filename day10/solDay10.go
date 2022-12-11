package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	data, _ := os.ReadFile("day10/inputDay10.txt")
	parsedData := parse(string(data))
	vals, crt := execute(parsedData)
	fmt.Println("Answer_10_1:", calculateVal(vals)) //13740
	fmt.Println("Answer_10_2:\n", render(crt))
}
func calculateVal(vals []int) int {
	var out int
	for _, v := range vals {
		out += v
	}
	return out
}
func draw(spritePos *int, cycle *int, multiplier *int) *string {
	cursor := *spritePos + *multiplier*40
	char := "."
	if (cursor == *cycle) || (cursor == *cycle-1) || (cursor == *cycle-2) {
		char = "#"
	}
	return &char
}
func render(out []*string) string {
	var render string
	for i := 0; i < len(out); i += 40 {
		line := out[i : i+40]
		var text string
		for _, v := range line {
			text = text + *v
		}
		render = fmt.Sprint(render, text, "\n ")
	}
	return render
}
func execute(input []Program) ([]int, []*string) {
	var signalStr []int
	InterestingCycle := 40 // checnge to 20 for sol 1
	cycle := 1
	value := 1
	out := make([]*string, 240)
	multiplier := 0
	for _, code := range input {
		if code.Instruction == "addx" {
			checkInteresting(&signalStr, &InterestingCycle, &cycle, &value, out, code, &multiplier)
			checkInteresting(&signalStr, &InterestingCycle, &cycle, &value, out, code, &multiplier)
			value += code.Value
		} else {
			checkInteresting(&signalStr, &InterestingCycle, &cycle, &value, out, code, &multiplier)
		}
	}
	return signalStr, out
}
func checkInteresting(signalStr *[]int, InterestingCycle *int, cycle *int, value *int, out []*string, code Program, multiplier *int) {
	fmt.Println("cycle ", *cycle, "cursor at ", *value, "will draw?", ((*value == *cycle) || (*value == *cycle-1) || (*value == *cycle-2)), "instruction ", code.Instruction, code.Value)
	out[*cycle-1] = draw(value, cycle, multiplier)
	if *cycle == *InterestingCycle {
		*InterestingCycle += 40
		*signalStr = append(*signalStr, *cycle*(*value)) //req for sol1
		*multiplier++
	}
	*cycle++
}
func parse(input string) []Program {
	var out []Program
	data := strings.Split(input, "\r\n")
	for _, v := range data {
		var temp Program
		x := strings.Split(v, " ")
		temp.Instruction = x[0]
		if strings.HasPrefix(x[0], "addx") {
			temp.Value, _ = strconv.Atoi(x[1])
		}
		out = append(out, temp)
	}
	return out
}

type Program struct {
	Instruction string
	Value       int
}
