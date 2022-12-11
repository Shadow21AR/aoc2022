package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/Knetic/govaluate.v2"
)

func Day11() {
	var monkeys []Monkey
	file, err := os.ReadFile("day11/inputDay11.txt")
	if err != nil {
		panic(err)
	}
	parseData(string(file), &monkeys)
	start(&monkeys)
	fmt.Println("Monkey business:", monkeyBusiness(monkeys))
}

func parseData(input string, monkeys *[]Monkey) {
	data := strings.Split(input, "\r\n\r\n")
	for _, monkey := range data {
		monkeyData := strings.Split(monkey, "\r\n ")
		createMonkeyData(monkeyData, monkeys)
	}
}
func createMonkeyData(input []string, monkeys *[]Monkey) {
	var newMonkey Monkey
	for _, v := range input {
		if strings.Contains(v, "Monkey") {
			newMonkey.Name, _ = strconv.Atoi(strings.Trim(strings.Split(v, " ")[1], ":"))
		}
		if strings.Contains(v, "Starting") {
			items := strings.Split(strings.SplitAfter(v, ": ")[1], ", ")
			for _, v := range items {
				n, _ := strconv.Atoi(v)
				newMonkey.Items = append(newMonkey.Items, n)
			}
		}
		if strings.Contains(v, "Operation") {
			operation := strings.Split(v, "= ")[1]
			newMonkey.Operation = operation
		}
		if strings.Contains(v, "Test") {
			operation := strings.Split(v, " ")
			newMonkey.TestCase, _ = strconv.Atoi(operation[len(operation)-1])
		}
		if strings.Contains(v, "throw") {
			passTo, _ := strconv.Atoi(strings.SplitAfter(v, "monkey ")[1])
			newMonkey.TestExec = append(newMonkey.TestExec, passTo)
		}
	}
	*monkeys = append(*monkeys, newMonkey)
}
func start(monkeys *[]Monkey) {
	var shitNum = 1
	for _, m := range *monkeys {
		shitNum *= m.TestCase
	}
	// for i := 1; i <= 20; i++ { // sol 1
	for i := 1; i <= 10000; i++ { // sol 2
		for n := range *monkeys {
			monkeyShit(monkeys, n, shitNum)
		}
	}
}
func monkeyShit(monkeys *[]Monkey, n int, shitNum int) {
	monkey := (*monkeys)[n]
	for _, v := range monkey.Items {
		// worry := operate(v, monkey.Operation) / 3 //sol1
		worry := operate(v, monkey.Operation) % shitNum //sol2
		toMonkey := monkey.TestExec[1]
		if worry%monkey.TestCase == 0 {
			toMonkey = monkey.TestExec[0]
		}
		(*monkeys)[toMonkey].Items = append((*monkeys)[toMonkey].Items, worry)
		monkey.Items = monkey.Items[1:]
		monkey.Inspected++
		(*monkeys)[n] = monkey
	}
}
func operate(worry int, operation string) int {
	expression, _ := govaluate.NewEvaluableExpression(operation)
	parameters := make(map[string]interface{}, 8)
	parameters["old"] = worry
	result, _ := expression.Evaluate(parameters)
	return int(result.(float64))
}
func monkeyBusiness(monkeys []Monkey) int {
	var inspected []int
	for _, v := range monkeys {
		inspected = append(inspected, v.Inspected)
	}
	sort.Ints(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

type Monkey struct {
	Name      int
	Items     []int
	Operation string
	TestCase  int
	TestExec  []int
	Inspected int
}
