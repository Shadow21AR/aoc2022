package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day2() {
	data, err := os.Open("day2/inputDay2.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var score int
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}
	//Part 1
	for _, v := range lines {
		moves := strings.Split(v, " ")
		switch {
		default:
			score = score + 3 + points[moves[1]]
		case v == "A Y" || v == "B Z" || v == "C X":
			score = score + 6 + points[moves[1]]
		case v == "A Z" || v == "B X" || v == "C Y":
			score = score + points[moves[1]]
		}
	}
	fmt.Println("Answer2_1: ", score)
	//Part 2
	var score2 int
	for _, v := range lines {
		moves := strings.Split(v, " ")
		switch moves[1] {
		case "X": //Gotta Lose
			score2 = score2 + points[lose(moves[0])]
		case "Y": //MEh Draw
			score2 = score2 + 3 + points[moves[0]]
		case "Z": //AYY Win
			score2 = score2 + 6 + points[win(moves[0])]
		}
	}
	fmt.Println("Answer2_2: ", score2)
}

func lose(elfMove string) string {
	switch elfMove {
	case "A": //rock
		return "Z"
	case "B": //paper
		return "X"
	case "C": //scissor
		return "Y"
	}
	return ""
}
func win(elfMove string) string {
	switch elfMove {
	case "A": //rock
		return "Y"
	case "B": //paper
		return "Z"
	case "C": //scissor
		return "A"
	}
	return ""
}
