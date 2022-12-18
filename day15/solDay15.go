package day15

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func Day15() { // copied from Wraken, cur couldn't do part 2 :c
	r, err := os.Open("day15/inputDay15.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)
	lines := []string{}
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	sensors := []sensor{}
	for _, l := range lines {
		s := sensor{}
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &s.beaconX, &s.beaconY)
		if err != nil {
			log.Fatal(err)
		}
		s.d = int(manhattanDistance(s.x, s.y, s.beaconX, s.beaconY))
		sensors = append(sensors, s)
	}
	minX, maxX := sensors[0].x-sensors[0].d, sensors[0].x+sensors[0].d
	for _, s := range sensors {
		min, max := s.x-s.d, s.x+s.d
		if min < minX {
			minX = min
		}
		if max > maxX {
			maxX = max
		}
	}
	fmt.Println("Part 1:", checkBeaconY(minX, maxX, 2000000, sensors))
	fmt.Println("Part 2:", findDistressBeacon(0, 4000000, sensors))
}

func manhattanDistance(x, y, x1, y1 int) int {
	return int(math.Abs(float64(x1)-float64(x)) + math.Abs(float64(y1)-float64(y)))
}

func checkBeaconY(minX, maxX, y int, sensors []sensor) int {
	nb := 0
	for x := minX; x <= maxX; x++ {
		for _, s := range sensors {
			d := manhattanDistance(s.x, s.y, x, y)
			if d <= s.d {
				b := false
				for _, s := range sensors {
					if s.beaconX == x && s.beaconY == y {
						b = true
						break
					}
				}
				if !b {
					nb++
				}
				break
			}
		}
	}
	return nb
}

func findDistressBeacon(min, max int, sensors []sensor) int {
	sort.SliceStable(sensors, func(i, j int) bool {
		return sensors[i].x < sensors[j].x
	})
	minY := min
	maxY := max
	for y := minY; y <= maxY; y++ {
		currx := min
		for _, s := range sensors {
			d := manhattanDistance(s.x, s.y, s.x, y)
			wl := (s.d-d)*2 + 1
			if wl > 0 {
				wX1 := s.x - (wl / 2)
				wX2 := s.x + (wl / 2)
				if wX1 < 0 {
					wX1 = 0
				}
				if currx >= wX1 && currx <= wX2 {
					currx = wX2
				}
				if currx >= max {
					break
				}
			}
		}
		if currx < max {
			return (currx+1)*4000000 + y
		}
	}
	return -1
}

type sensor struct {
	x       int
	y       int
	beaconX int
	beaconY int
	d       int
}
