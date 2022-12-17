package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var test int

func Day14() {
	caveMap := map[Pos]Tile{}
	file, _ := os.ReadFile("day14/inputDay14.txt")
	makeMap(string(file), &caveMap)
	caveMap[Pos{500, 0}] = Tile{X: 500, Y: 0}
	source := (caveMap)[Pos{500, 0}]
	keys := make([][]int, 0)
	for key := range caveMap {
		keys = append(keys, []int{key.X, key.Y})
	}
	fillAir(&caveMap, keys)
	simulate(&caveMap, source, keys) //part1
	caveMap2 := make(map[Pos]Tile)
	for k, v := range caveMap {
		caveMap2[k] = v
	}
	part2(&caveMap2, keys) //part2
}

func simulate(caveMap *map[Pos]Tile, curPos Tile, keys [][]int) {
	// time.Sleep(2500 * time.Microsecond) // to see better lmao
	downPos := (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
	leftPos := (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
	rightPos := (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
	if downPos.Air {
		curPos = downPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
		simulate(caveMap, curPos, keys)
		return
	} else if leftPos.Air {
		curPos = leftPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
		simulate(caveMap, curPos, keys)
		return
	} else if rightPos.Air {
		curPos = rightPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
		simulate(caveMap, curPos, keys)
		return
	} else if !downPos.Exist && !leftPos.Exist && !rightPos.Exist {
		countSand(caveMap, keys)
		return
	} else {
		curPos.Sand = true
		curPos.Air = false
		(*caveMap)[Pos{curPos.X, curPos.Y}] = curPos
		source := (*caveMap)[Pos{500, 0}]
		simulate(caveMap, source, keys)
		return
	}
}

func countSand(caveMap *map[Pos]Tile, keys [][]int) {
	min, max := getMinMax(keys)
	var sand int
	for i := min[0]; i <= max[0]; i++ {
		for k := min[1]; k <= max[1]; k++ {
			if (*caveMap)[Pos{i, k}].Sand {
				sand++
			}
		}
	}
	fmt.Println("Answer ", sand)
}

func makeMap(input string, caveMap *map[Pos]Tile) {
	for _, v := range strings.Split(input, "\r\n") {
		createRock(strings.Split(v, " -> "), caveMap)
	}
}

func createRock(input []string, caveMap *map[Pos]Tile) {
	start := createTile(strings.Split(input[0], ","), caveMap, true)
	end := createTile(strings.Split(input[1], ","), caveMap, true)
	fillRock(start, end, caveMap)
	if len(input) > 2 {
		createRock(input[1:], caveMap)
	}
}

func createTile(cord []string, caveMap *map[Pos]Tile, rock bool) *Tile {
	x, _ := strconv.Atoi(cord[0])
	y, _ := strconv.Atoi(cord[1])
	tilePos := Pos{}
	tilePos.X = x
	tilePos.Y = y
	newTile := Tile{}
	newTile.X = x
	newTile.Y = y
	newTile.Rock = rock
	newTile.Air = !rock
	newTile.Exist = true
	(*caveMap)[tilePos] = newTile
	return &newTile
}

func fillRock(start *Tile, end *Tile, caveMap *map[Pos]Tile) {
	if start.X > end.X {
		for i := 1; i <= start.X-end.X; i++ {
			newRock := []string{strconv.Itoa(start.X - i), strconv.Itoa(start.Y)}
			createTile(newRock, caveMap, true)
		}
	} else if start.X < end.X {
		for i := 1; i <= end.X-start.X; i++ {
			newRock := []string{strconv.Itoa(start.X + i), strconv.Itoa(start.Y)}
			createTile(newRock, caveMap, true)
		}
	} else {
		if start.Y > end.Y {
			for i := 1; i <= start.Y-end.Y; i++ {
				newRock := []string{strconv.Itoa(start.X), strconv.Itoa(start.Y - i)}
				createTile(newRock, caveMap, true)
			}
		} else if start.Y < end.Y {
			for i := 1; i <= end.Y-start.Y; i++ {
				newRock := []string{strconv.Itoa(start.X), strconv.Itoa(start.Y + i)}
				createTile(newRock, caveMap, true)
			}
		}
	}
}

func getMinMax(keys [][]int) ([]int, []int) {
	min := []int{500, 500}
	max := []int{0, 0}
	for _, key := range keys {
		if key[0] > max[0] {
			max[0] = key[0]
		}
		if key[1] > max[1] {
			max[1] = key[1]
		}
		if key[0] < min[0] {
			min[0] = key[0]
		}
		if key[1] < min[1] {
			min[1] = key[1]
		}
	}
	return min, max
}

func fillAir(caveMap *map[Pos]Tile, keys [][]int) {
	min, max := getMinMax(keys)
	for i := min[0]; i <= max[0]; i++ {
		for k := min[1]; k <= max[1]; k++ {
			if !(*caveMap)[Pos{i, k}].Exist {
				newRock := []string{strconv.Itoa(i), strconv.Itoa(k)}
				createTile(newRock, caveMap, false)
			}
		}
	}
}

func part2(caveMap *map[Pos]Tile, keys [][]int) {
	min, max := getMinMax(keys)
	minRock := []string{strconv.Itoa(min[0] - max[1]), strconv.Itoa(max[1] + 2)}
	maxRock := []string{strconv.Itoa(max[0] + max[1]), strconv.Itoa(max[1] + 2)}
	start := createTile(minRock, caveMap, true)
	end := createTile(maxRock, caveMap, true)
	fillRock(start, end, caveMap)
	keys = make([][]int, 0)
	for key := range *caveMap {
		keys = append(keys, []int{key.X, key.Y})
	}
	fillAir(caveMap, keys)
	curPos := (*caveMap)[Pos{500, 0}]
	curPos = simulate2(caveMap, curPos, keys)
	for {
		if !(curPos.X == 0 && curPos.Y == 0) {
			curPos = simulate2(caveMap, curPos, keys)
		} else {
			break
		}
	}
}
func simulate2(caveMap *map[Pos]Tile, curPos Tile, keys [][]int) Tile {
	downPos := (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
	leftPos := (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
	rightPos := (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
	if downPos.Air {
		curPos = downPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
	} else if leftPos.Air {
		curPos = leftPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
	} else if rightPos.Air {
		curPos = rightPos
		downPos = (*caveMap)[Pos{curPos.X, curPos.Y + 1}]
		leftPos = (*caveMap)[Pos{curPos.X - 1, curPos.Y + 1}]
		rightPos = (*caveMap)[Pos{curPos.X + 1, curPos.Y + 1}]
	} else {
		if curPos.X == 500 && curPos.Y == 0 {
			curPos.Sand = true
			curPos.Air = false
			(*caveMap)[Pos{curPos.X, curPos.Y}] = curPos
			countSand(caveMap, keys)
			return Tile{}
		}
		curPos.Sand = true
		curPos.Air = false
		(*caveMap)[Pos{curPos.X, curPos.Y}] = curPos
		curPos = (*caveMap)[Pos{500, 0}]
	}
	return curPos
}

func render(caveMap *map[Pos]Tile, keys [][]int) { //it was fun, not using now xD
	min, max := getMinMax(keys)
	out := make([][]string, 0)
	for i := min[0]; i <= max[0]; i++ {
		temp := make([]string, 0)
		for k := min[1]; k <= max[1]; k++ {
			if (*caveMap)[Pos{i, k}].Rock {
				temp = append(temp, "X")
			} else if (*caveMap)[Pos{i, k}].Sand {
				temp = append(temp, "\u001b[2;33mO\u001b[0m") //run on linux only (or whatever)
				// temp = append(temp, "\u001b[2;33mO\u001b[0m") //for windows
			} else {
				temp = append(temp, " ")
			}
		}
		out = append(out, temp)
	}
	out[500-min[0]][0] = "+"
	outPrint(out)
}

func outPrint(out [][]string) {
	data := ""
	for i := range out[0] {
		temp := ""
		for _, v := range out {
			temp = temp + v[i]
		}
		data = data + temp + "\n"
	}
	fmt.Printf("%s", data)
}

type Tile struct {
	X     int
	Y     int
	Rock  bool
	Air   bool
	Sand  bool
	Exist bool
}
type Pos struct {
	X int
	Y int
}
