package day16

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day16() {
	valves := map[string]int{}
	tunnels := map[string][]string{}
	dists := map[string][]Dists{}
	relevant := []string{}
	index := map[string]int{}
	file, _ := os.ReadFile("day16/inputDay16.txt")
	parseData(string(file), &valves, &tunnels, &dists, &relevant)
	for i, val := range relevant {
		index[val] = i
	}
	// fmt.Println("Answer 16_1", dfs(&valves, &dists, 30, "AA", 0, index)) //sol1
	max := 0
	maxBit := 1<<len(relevant) - 1
	for i := 0; i <= (maxBit+1)/2; i++ {
		x := dfs(&valves, &dists, 26, "AA", i, index) + dfs(&valves, &dists, 26, "AA", maxBit^i, index)
		max = int(math.Max(float64(max), float64(x)))
	}
	fmt.Println("Answer 16_2", max) //sol2
}

func parseData(input string, valves *map[string]int, tunnels *map[string][]string, dists *map[string][]Dists, relevant *[]string) {
	for _, v := range strings.Split(input, "\r\n") {
		data := strings.Split(v, " ")
		(*valves)[data[1]] = toInt(strings.Split(strings.Split(data[4], "=")[1], ";")[0])
		(*tunnels)[data[1]] = strings.Split(strings.Join(data[9:], ""), ",")
	}
	for valve := range *valves {
		if !(valve != "AA" && (*valves)[valve] <= 0) {
			*relevant = append(*relevant, valve)
			(*dists)[valve] = []Dists{}
			visited := map[string]int{}
			queue := []Queue{{Dist: 0, Pos: valve}}
			for len(queue) > 0 {
				pos := queue[0]
				queue = queue[1:]
				for _, neighbor := range (*tunnels)[pos.Pos] {
					if checkVisited(neighbor, &visited) {
						continue
					}
					visited[neighbor] = 1
					if neighbor != valve {
						if (*valves)[neighbor] > 0 {
							(*dists)[valve] = append((*dists)[valve], Dists{Valve: neighbor, Distance: pos.Dist + 1})
						}
						queue = append(queue, Queue{Dist: pos.Dist + 1, Pos: neighbor})
					}
				}
			}
		}
	}
}

func dfs(valves *map[string]int, dists *map[string][]Dists, time int, curVal string, bitmask int, index map[string]int) (out int) {
	cache := map[Cache]int{}
	v, ok := cache[Cache{time, curVal, bitmask}]
	if ok {
		return v
	}
	out = 0
	for _, dist := range (*dists)[curVal] {
		bit := 1 << index[dist.Valve]
		if bitmask&bit > 0 {
			continue
		}
		timeLeft := time - dist.Distance - 1
		if timeLeft > 0 {
			x := dfs(valves, dists, timeLeft, dist.Valve, bitmask|bit, index) + (*valves)[dist.Valve]*timeLeft
			out = int(math.Max(float64(out), float64(x)))
		}
	}
	cache[Cache{time, curVal, bitmask}] = out
	return out
}

func toInt(input string) (out int) {
	out, _ = strconv.Atoi(input)
	return
}
func checkVisited(neighbor string, visited *map[string]int) bool {
	_, ok := (*visited)[neighbor]
	return ok
}

type Dists struct {
	Valve    string
	Distance int
}
type Queue struct {
	Dist int
	Pos  string
}
type Cache struct {
	x int
	y string
	z int
}
