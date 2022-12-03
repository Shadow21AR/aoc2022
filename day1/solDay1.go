package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	data, err := os.Open("day1/inputDay1.txt")
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
	var cal []int
	var temp int
	for _, v := range lines {
		num, err := strconv.Atoi(v)
		temp = temp + num
		if err != nil {
			cal = append(cal, temp)
			temp = 0
		}
	}
	sort.Ints(cal)
	fmt.Println("Answer1 = ", cal[len(cal)-1])
	fmt.Println("Answer2 = ", (cal[len(cal)-1] + cal[len(cal)-2] + cal[len(cal)-3]))
}
