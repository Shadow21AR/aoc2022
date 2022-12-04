package main

import (
	"fmt"
	"log"
	"os"

	"aoc2k22/shadow/day1"
	"aoc2k22/shadow/day2"
	"aoc2k22/shadow/day3"
	"aoc2k22/shadow/day4"
)

func main() {
	var day int
	fmt.Println("Which Day: ")
	fmt.Scan(&day)

	type Day func(data *os.File)

	Days := map[int]Day{
		1: day1.Day1,
		2: day2.Day2,
		3: day3.Day3,
		4: day4.Day4,
	}

	log.Printf("Executing AOC 2022 Day %d", day)
	input, err := ReadFile(fmt.Sprintf("./day%d/inputDay%[1]d.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	Days[day](input)

}

func ReadFile(loc string) (*os.File, error) {
	data, err := os.Open(loc)
	if err != nil {
		panic(err)
	}
	return data, nil
}
