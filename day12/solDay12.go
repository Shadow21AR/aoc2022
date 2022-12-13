package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day12() {
	var heightMap [][]rune
	nodes := make(map[string]Node, 0)
	queue := make([]string, 0)
	file, err := os.ReadFile("day12/inputDay12.txt")
	if err != nil {
		panic(err)
	}
	parseInput(string(file), &heightMap)
	// nodes[findIt(heightMap, 'S')] = Node{}                  //sol1
	// bfs(&heightMap, &nodes, &queue, findIt(heightMap, 'S')) //sol1
	nodes[findIt(heightMap, 'E')] = Node{}                  //sol2
	bfs(&heightMap, &nodes, &queue, findIt(heightMap, 'E')) //sol2
}

func bfs(heightMap *[][]rune, nodes *map[string]Node, queue *[]string, cNode string) {
	var curPos []int
	curNodePos := strings.Split(cNode, " ")
	for _, v := range curNodePos {
		n, _ := strconv.Atoi(v)
		curPos = append(curPos, n)
	}
	curNode := (*nodes)[cNode]
	curNode.Name = cNode
	curNode.Visited = true
	curNode.Value = (*heightMap)[curPos[0]][curPos[1]]
	(*nodes)[cNode] = curNode
	if curPos[0] > 0 {
		newName := fmt.Sprint(curPos[0]-1, curPos[1])
		if makeNewNode(heightMap, nodes, &curNode, newName, queue, curPos) {
			return
		}
	}
	if curPos[0] < len(*heightMap)-1 {
		newName := fmt.Sprint(curPos[0]+1, curPos[1])
		if makeNewNode(heightMap, nodes, &curNode, newName, queue, curPos) {
			return
		}
	}
	if curPos[1] > 0 {
		newName := fmt.Sprint(curPos[0], curPos[1]-1)
		if makeNewNode(heightMap, nodes, &curNode, newName, queue, curPos) {
			return
		}
	}
	if curPos[1] < len((*heightMap)[0])-1 {
		newName := fmt.Sprint(curPos[0], curPos[1]+1)
		if makeNewNode(heightMap, nodes, &curNode, newName, queue, curPos) {
			return
		}
	}
	if len(*queue) > 0 {
		nextNode := (*queue)[0]
		*queue = (*queue)[1:]
		bfs(heightMap, nodes, queue, nextNode)
	}
}

func makeNewNode(heightMap *[][]rune, nodes *map[string]Node, curNode *Node, newName string, queue *[]string, curPos []int) (stop bool) {
	_, ok := (*nodes)[newName]
	if !ok {
		(*nodes)[newName] = Node{}
	}
	var newPos []int
	newNodePos := strings.Split(newName, " ")
	for _, v := range newNodePos {
		n, _ := strconv.Atoi(v)
		newPos = append(newPos, n)
	}
	newNode := (*nodes)[newName]
	newNode.Value = (*heightMap)[newPos[0]][newPos[1]]
	// if curNode.Value == 'S' { //sol1
	// 	curNode.Value = 'a'
	// }
	if curNode.Value == 'E' { //sol2
		curNode.Value = 'z'
	}
	// if !newNode.Visited && newNode.Value-curNode.Value < 2 { //sol1
	// 	if !(newNode.Value == 'E' && curNode.Value != 'z') { //sol1
	if !newNode.Visited && newNode.Value-curNode.Value > -2 { //sol2
		if !(newNode.Value == 'a' && curNode.Value != 'b') { //sol2
			newNode.Name = newName
			newNode.Parent = curNode
			newNode.HasParent = true
			curNode.Child = append(curNode.Child, newName)
			newNode.Visited = true
			(*nodes)[newName] = newNode
			(*nodes)[curNode.Name] = *curNode
			*queue = append(*queue, newName)
			// if (*nodes)[newName].Value == 'E' { //sol1
			if (*nodes)[newName].Value == 'a' { //sol2
				fmt.Println("here")
				fmt.Println(findSteps(newNode))
				stop = true
			}
		}
	}
	return stop
}

func findIt(heightMap [][]rune, des rune) (out string) {
	for i, v := range heightMap {
		for n, r := range v {
			if r == des {
				out = fmt.Sprint(i, n)
			}
		}
	}
	return out
}
func parseInput(file string, heightMap *[][]rune) {
	lines := strings.Split(file, "\r\n")
	for _, v := range lines {
		r := []rune(v)
		*heightMap = append(*heightMap, r)
	}
}

// S 83, E 69
func findSteps(node Node) (out int) {
	for node.HasParent {
		node = *node.Parent
		out++
	}
	return out
}

type Node struct {
	Name      string
	Value     rune
	Parent    *Node
	HasParent bool
	Child     []string
	Visited   bool
}
