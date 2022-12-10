package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day8() {
	file, err := os.ReadFile("day8/inputDay8.txt")
	if err != nil {
		panic(err)
	}
	forest := makeTrees(string(file))
	visible, scene := findVisible(forest)
	visibleTrees := visible + len(forest)*2 + (len(forest[0])-2)*2
	fmt.Println("Answer 8_1:", visibleTrees, "\nAnswer 8_2:", scene)
}
func makeTrees(file string) [][]int {
	trees := make([][]int, 0)
	rows := strings.Split(file, "\n")
	for _, v := range rows {
		newRow := strings.Split(v, "")
		temp := make([]int, 0)
		for i := range newRow {
			num, err := strconv.Atoi(newRow[i])
			if err == nil {
				temp = append(temp, num)
			}
		}
		trees = append(trees, temp)
	}
	return trees
}

func findVisible(forest [][]int) (int, int) {
	trees := forest[1 : len(forest[0])-1]
	visible := 0
	scene := 0
	for i := range trees {
		treerow := trees[i][1 : len(trees[i])-1]
		for ii := range treerow {
			v, s := checkVisibility(forest, i, ii)
			visible += v
			if s > scene {
				scene = s
			}
		}
	}
	return visible, scene
}

func checkVisibility(forest [][]int, i int, ii int) (int, int) {
	bot, bots := checkBot(forest, i, ii)
	top, tops := checkTop(forest, i, ii)
	left, lefts := checkLeft(forest, i, ii)
	right, rights := checkRight(forest, i, ii)
	if bot || top || left || right {
		return 1, (bots * tops * lefts * rights)
	}
	return 0, (bots * tops * lefts * rights)
}

func checkTop(forest [][]int, i int, ii int) (bool, int) {
	scene := 0
	currentTree := forest[i+1][ii+1]
	for _, y := range reverseShit2(forest[:i+1]) {
		scene++
		if currentTree <= y[ii+1] {
			return false, scene
		}
	}
	return true, scene
}

func checkBot(forest [][]int, i int, ii int) (bool, int) {
	scene := 0
	currentTree := forest[i+1][ii+1]
	for _, y := range forest[i+2:] {
		scene++
		if currentTree <= y[ii+1] {
			return false, scene
		}
	}
	return true, scene
}

func checkLeft(forest [][]int, i int, ii int) (bool, int) {
	scene := 0
	currentTree := forest[i+1][ii+1]
	for _, y := range reverseShit(forest[i+1][:ii+1]) {
		scene++
		if currentTree <= y {
			return false, scene
		}
	}
	return true, scene
}

func checkRight(forest [][]int, i int, ii int) (bool, int) {
	scene := 0
	currentTree := forest[i+1][ii+1]
	for _, y := range forest[i+1][ii+2:] {
		scene++
		if currentTree <= y {
			return false, scene
		}
	}
	return true, scene
}
func reverseShit(things []int) (out []int) {
	for i := range things {
		out = append(out, things[len(things)-i-1])
	}
	return out
}
func reverseShit2(things [][]int) (out [][]int) {
	for i := range things {
		out = append(out, things[len(things)-i-1])
	}
	return out
}
