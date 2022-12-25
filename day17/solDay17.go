package day17

import (
	"image"
	"log"
	"os"
	"strings"
)

var rocksCount int

func Day17() {
	jets := []int{}
	chamber := make([][]string, 0)
	var curRock, curJet int
	rocks := []Rock{
		{[]Pos{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		{[]Pos{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}},
		{[]Pos{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}},
		{[]Pos{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		{[]Pos{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},
	}
	file, _ := os.ReadFile("day17/inputDay17.txt")
	for _, v := range strings.Split(string(file), "") {
		if v == ">" {
			jets = append(jets, 1)
		} else {
			jets = append(jets, -1)
		}
	}
	spawnRock(&chamber, rocks, &curRock, jets, &curJet) //part 1
	part2()                                             //part 2 got it from someone else
}

func part2() {
	input, _ := os.ReadFile("day17/inputDay17.txt")
	jets := strings.TrimSpace(string(input))
	rocks := [][]image.Point{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}},
		{{2, 2}, {2, 1}, {0, 0}, {1, 0}, {2, 0}},
		{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		{{0, 1}, {1, 1}, {0, 0}, {1, 0}},
	}
	grid := map[image.Point]struct{}{}
	move := func(rock []image.Point, delta image.Point) bool {
		nrock := make([]image.Point, len(rock))
		for i, p := range rock {
			p = p.Add(delta)
			if _, ok := grid[p]; ok || p.X < 0 || p.X >= 7 || p.Y < 0 {
				return false
			}
			nrock[i] = p
		}
		copy(rock, nrock)
		return true
	}
	cache := map[[2]int][]int{}
	height, jet := 0, 0
	for i := 0; i < 1000000000000; i++ {
		k := [2]int{i % len(rocks), jet}
		if c, ok := cache[k]; ok {
			if n, d := 1000000000000-i, i-c[0]; n%d == 0 {
				log.Println(height + n/d*(height-c[1]))
				break
			}
		}
		cache[k] = []int{i, height}
		rock := []image.Point{}
		for _, p := range rocks[i%len(rocks)] {
			rock = append(rock, p.Add(image.Point{2, height + 3}))
		}
		for {
			move(rock, image.Point{int(jets[jet]) - int('='), 0})
			jet = (jet + 1) % len(jets)
			if !move(rock, image.Point{0, -1}) {
				for _, p := range rock {
					grid[p] = struct{}{}
					if p.Y+1 > height {
						height = p.Y + 1
					}
				}
				break
			}
		}
	}
}

func spawnRock(chamber *[][]string, rocks []Rock, curRock *int, jets []int, curJet *int) {
	rocksCount++
	if rocksCount == 2022+1 {
		countHeight(*chamber)
		return
	}
	rock := Rock{}
	temp := rocks[*curRock]
	rock.Shape = append(rock.Shape, temp.Shape...)
	*curRock++
	if *curRock >= len(rocks) {
		*curRock = 0
	}
	left := 2
	bot := 3
	for i, v := range *chamber {
		if strings.Join([]string{" ", " ", " ", " ", " ", " ", " "}, "") == strings.Join(v, "") {
			bot = i + 3
			break
		}
	}
	for i := range rock.Shape {
		rock.Shape[i].X += left
		rock.Shape[i].Y += bot
	}
	if len(*chamber) <= bot+3 {
		makeChamber(chamber, bot)
	}
	for _, v := range rock.Shape {
		(*chamber)[v.Y][v.X] = "@"
	}
	playRock(chamber, jets, curJet, &rock, rocks, curRock)
}
func playRock(chamber *[][]string, jets []int, curJet *int, movingRock *Rock, rocks []Rock, curRock *int) {
	for {
		jet := jets[*curJet]
		moveRockWithJet(chamber, movingRock, jet)
		*curJet++
		if *curJet >= len(jets) {
			*curJet = 0
		}
		for _, v := range movingRock.Shape {
			if v.Y == 0 || (*chamber)[v.Y-1][v.X] == "X" {
				newFunctionNameLmao(chamber, *movingRock)
				spawnRock(chamber, rocks, curRock, jets, curJet)
				return
			}
		}
		moverRockDown(chamber, movingRock)
	}
}

func moveRockWithJet(chamber *[][]string, rock *Rock, jet int) {
	newRock := Rock{}
	left := 6
	right := 0
	for _, v := range rock.Shape {
		if v.X < left {
			left = v.X
		}
		if v.X > right {
			right = v.X
		}
	}
	if jet+right > 6 || jet+left < 0 {
		return
	}
	for _, v := range rock.Shape {
		if (*chamber)[v.Y][v.X+jet] == "X" {
			return
		}
	}
	for _, v := range rock.Shape {
		(*chamber)[v.Y][v.X] = " "
	}
	for _, v := range rock.Shape {
		(*chamber)[v.Y][v.X+jet] = "@"
		newRock.Shape = append(newRock.Shape, Pos{v.X + jet, v.Y})
	}
	*rock = newRock
}

func moverRockDown(chamber *[][]string, movingRock *Rock) {
	newRock := Rock{}
	for _, v := range movingRock.Shape {
		(*chamber)[v.Y][v.X] = " "
	}
	for _, v := range movingRock.Shape {
		(*chamber)[v.Y-1][v.X] = "@"
		newRock.Shape = append(newRock.Shape, Pos{v.X, v.Y - 1})
	}
	*movingRock = newRock
}

func newFunctionNameLmao(chamber *[][]string, movingRock Rock) {
	for _, v := range movingRock.Shape {
		(*chamber)[v.Y][v.X] = "X"
	}
}

func makeChamber(chamber *[][]string, bot int) {
	for i := 0; i <= bot; i++ {
		temp := []string{" ", " ", " ", " ", " ", " ", " "}
		*chamber = append(*chamber, temp)
	}
}

func countHeight(chamber [][]string) {
	var h int
	for _, v := range chamber {
		if !(strings.Join(v, "") == "       ") {
			h++
		}
	}
	log.Println(h)
}

type Rock struct {
	Shape []Pos
}
type Pos struct {
	X, Y int
}
