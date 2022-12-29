package day20

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// copied it from dhruvmanila, cuz i gave up :< see failed attempt :<

const decryptionKey = 811589153

type positionKey struct {
	index int
	value int
}

func constructList(numbers []int) (map[positionKey]*ring.Ring, positionKey) {
	positions := make(map[positionKey]*ring.Ring, len(numbers))
	r := ring.New(len(numbers))
	zeroKey := positionKey{value: 0}
	for idx, number := range numbers {
		if number == 0 {
			zeroKey.index = idx
		}
		positions[positionKey{idx, number}] = r
		r.Value = number
		r = r.Next()
	}
	return positions, zeroKey
}
func mix(numbers []int, n int) (coordinateSum int) {
	positions, zeroKey := constructList(numbers)

	length := len(numbers) - 1
	halflen := length >> 1
	for ; n > 0; n-- {
		for idx, number := range numbers {
			r := positions[positionKey{idx, number}].Prev()
			removed := r.Unlink(1)

			if (number > halflen) || (number < -halflen) {
				number %= length
				switch {
				case number > halflen:
					number -= length
				case number < -halflen:
					number += length
				}
			}
			r.Move(number).Link(removed)
		}
	}
	r := positions[zeroKey]
	for i := 1; i <= 3; i++ {
		r = r.Move(1000)
		coordinateSum += r.Value.(int)
	}
	return coordinateSum
}

func toInt(in []string) []int {
	out := []int{}
	for _, v := range in {
		n, _ := strconv.Atoi(v)
		out = append(out, n)
	}
	return out
}

func Day20() {
	file, _ := os.ReadFile("day20/inputDay20.txt")
	numbers := toInt(strings.Split(string(file), "\r\n"))
	coordinateSum1 := mix(numbers, 1)

	for idx := range numbers {
		numbers[idx] *= decryptionKey
	}
	coordinateSum2 := mix(numbers, 10)
	fmt.Printf("20.1: %d\n20.2: %d\n", coordinateSum1, coordinateSum2)
}
