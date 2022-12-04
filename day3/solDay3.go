package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3(data *os.File) {
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var chars []int
	for _, v := range lines {
		len := strconv.Itoa(len(v) / 2)
		re := regexp.MustCompile(`(^\w{` + regexp.QuoteMeta(len) + `})(\w{` + regexp.QuoteMeta(len) + `}$)`)
		list := re.FindAllStringSubmatch(v, -1)
		for _, items := range list {
			for _, char1 := range items[1] {
				if strings.ContainsRune(items[2], char1) {
					chars = append(chars, int(char1))
					break
				}
			}
		}
	}
	fmt.Println("Priority is :", priority(chars))
	Day3_2()
}
func priority(chars []int) int {
	var priority int
	for _, rune := range chars {
		if rune > 97 {
			priority = priority + (rune - 96)
		} else {
			priority = priority + (rune - 38)
		}
	}
	return priority
}

func Day3_2() {
	data, err := os.Open("day3/inputDay3.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var tempLines []string
	var lines [][]string
	var temp int
	for scanner.Scan() {
		tempLines = append(tempLines, scanner.Text())
		temp++
		if temp == 3 {
			lines = append(lines, tempLines)
			tempLines = []string{}
			temp = 0
		}
	}
	var chars []int
	for _, v := range lines {
		freqMap := make([]map[rune]int, 0)
		for _, s := range v {
			tempy := make(map[rune]int)
			for _, r := range s {
				tempy[r] = 1
			}
			freqMap = append(freqMap, tempy)
		}
		temp2 := make(map[rune]int)
		for _, v := range freqMap {
			for key, val := range v {
				temp2[key] = temp2[key] + val
			}
		}
		for k, v := range temp2 {
			if v == 3 {
				chars = append(chars, int(k))
			}
		}
	}
	fmt.Println("Answer3_2 :", priority(chars))
}

// Priority is : 7811
// Answer3_2 : 2639
