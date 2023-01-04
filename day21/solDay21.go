package day21

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/Knetic/govaluate.v2"
)

func Day21() {
	t := time.Now()
	monkeys := map[string]Monkeys{}
	queue := []string{}
	file, err := os.ReadFile("day21/inputDay21.txt")
	if err != nil {
		log.Panic(err)
	}
	for _, v := range strings.Split(string(file), "\r\n") {
		data := strings.Split(v, ": ")
		x := regexp.MustCompile(`\d+`).FindAllString(data[1], -1)
		if len(x) > 0 {
			monkeys[data[0]] = Monkeys{Data: toInt(x[0])}
		} else {
			temp := strings.Split(data[1], " ")
			monkeys[data[0]] = Monkeys{Operation: temp[1], Others: []string{temp[0], temp[2]}}
			queue = append(queue, data[0])
		}
	}
	for {
		if len(queue) > 0 {
			monkey := monkeys[queue[0]]
			if monkeys[monkey.Others[0]].Data > 0 && monkeys[monkey.Others[1]].Data > 0 {
				monkey.Data = evaluate([]int{monkeys[monkey.Others[0]].Data, monkeys[monkey.Others[1]].Data}, monkey.Operation)
				monkeys[queue[0]] = monkey
				queue = queue[1:]
			} else {
				queue = append(queue[1:], queue[0])
			}
		} else {
			break
		}
	}
	log.Println("Answer 21_1: ", monkeys["root"].Data, ",took", time.Since(t)) //158661812617812
	nT := time.Now()
	for k, v := range monkeys {
		if len(v.Others) > 0 {
			x := monkeys[v.Others[0]]
			x.Parent = k
			monkeys[v.Others[0]] = x
			y := monkeys[v.Others[1]]
			y.Parent = k
			monkeys[v.Others[1]] = y
		}
	}
	newQueue := []string{}
	current := monkeys["humn"].Parent
	for current != "root" {
		newQueue = append(newQueue, current)
		current = monkeys[current].Parent
	}
	var total int
	root := monkeys["root"]
	if root.Others[0] != newQueue[len(newQueue)-1] {
		total = monkeys[root.Others[0]].Data
	} else {
		total = monkeys[root.Others[1]].Data
	}
	for i := len(newQueue) - 1; i >= 0; i-- {
		monkey := monkeys[newQueue[i]]
		if (i == 0 && monkey.Others[0] != "humn") || (i > 0 && monkey.Others[0] != newQueue[i-1]) {
			total = eval(monkeys[monkey.Others[0]].Data, total, monkey.Operation, true)
		} else {
			total = eval(monkeys[monkey.Others[1]].Data, total, monkey.Operation, false)
		}
	}
	log.Println("Answer 21_2: ", total, ",took", time.Since(nT), "\nTotal time taken:", time.Since(t)) //3352886133831
}

func eval(x, y int, op string, pos bool) int { //had to check somewhere, here and there
	var out int
	if op == "+" {
		out = y - x
	} else if op == "-" {
		if pos {
			out = (y - x) / -1
		} else {
			out = y + x
		}
	} else if op == "*" {
		out = y / x
	} else if op == "/" {
		if pos {
			out = x / y
		} else {
			out = y * x
		}
	}
	return out
}

func evaluate(x []int, op string) int {
	expression, _ := govaluate.NewEvaluableExpression("x" + op + "y")
	result, _ := expression.Evaluate(map[string]interface{}{"x": x[0], "y": x[1]})
	return int(result.(float64))
}

// func pwint(in map[string]Monkeys) {
// 	for k, v := range in {
// 		log.Printf("%s - %s %s = %d", k, v.Others, v.Operation, v.Data)
// 	}
// }

func toInt(in string) int {
	out, _ := strconv.Atoi(in)
	return out
}

type Monkeys struct {
	Others    []string
	Operation string
	Data      int
	Parent    string
}
