package day18

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	cubes := map[Cube]Cube{}
	file, _ := os.ReadFile("day18/inputDay18.txt")
	for _, v := range strings.Split(string(file), "\r\n") {
		cube := makeCube(strings.Split(v, ","))
		cubes[cube] = cube
	}
	log.Println("Answer 18_1:", part1(cubes))
	log.Println("Answer 18_2:", part2(cubes))
}

func part1(cubes map[Cube]Cube) int {
	var temp int
	for _, x := range cubes {
		neighbors := []Cube{{x.X + 1, x.Y, x.Z}, {x.X - 1, x.Y, x.Z}, {x.X, x.Y + 1, x.Z}, {x.X, x.Y - 1, x.Z}, {x.X, x.Y, x.Z + 1}, {x.X, x.Y, x.Z - 1}}
		for _, n := range neighbors {
			if !checkCube(n, cubes) {
				temp++
			}
		}
	}
	return temp
}

func part2(cubes map[Cube]Cube) int {
	cubeMap := buildMap(cubes)
	cubeMap.dfs(Cube{0, 0, 0})
	return cubeMap.area()
}

func buildMap(cubes map[Cube]Cube) cubeMap {
	myMap := [25][25][25]Node{}
	for _, v := range cubes {
		myMap[v.X+1][v.Y+1][v.Z+1].isCube = true
	}
	return myMap
}

func (m *cubeMap) dfs(x Cube) {
	m[x.X][x.Y][x.Z].Visited = true
	lenM := len(m)
	neighbors := []Cube{{x.X + 1, x.Y, x.Z}, {x.X - 1, x.Y, x.Z}, {x.X, x.Y + 1, x.Z}, {x.X, x.Y - 1, x.Z}, {x.X, x.Y, x.Z + 1}, {x.X, x.Y, x.Z - 1}}
	for _, p := range neighbors {
		if p.X < 0 || p.X >= lenM || p.Y < 0 || p.Y >= lenM || p.Z < 0 || p.Z >= lenM {
			continue
		}
		if !m[p.X][p.Y][p.Z].isCube && !m[p.X][p.Y][p.Z].Visited {
			m.dfs(p)
		}
	}
}

func (m *cubeMap) area() int {
	area := 0
	L := len(m)
	for x := range m {
		for y := range m[x] {
			for z := range m[x][y] {
				if !m[x][y][z].isCube {
					continue
				}
				neighbors := []Cube{{x + 1, y, z}, {x - 1, y, z}, {x, y + 1, z}, {x, y - 1, z}, {x, y, z + 1}, {x, y, z - 1}}

				for _, p := range neighbors {
					if p.X < 0 || p.X >= L || p.Y < 0 || p.Y >= L || p.Z < 0 || p.Z >= L {
						continue
					}
					if m[p.X][p.Y][p.Z].Visited {
						area += 1
					}
				}
			}
		}
	}
	return area
}

func checkCube(c Cube, cubes map[Cube]Cube) bool {
	_, ok := cubes[c]
	return ok
}

func makeCube(input []string) (out Cube) {
	out.X, _ = strconv.Atoi(input[0])
	out.Y, _ = strconv.Atoi(input[1])
	out.Z, _ = strconv.Atoi(input[2])
	return out
}

type Cube struct {
	X, Y, Z int
}
type Node struct {
	Visited, isCube bool
}
type cubeMap [25][25][25]Node
