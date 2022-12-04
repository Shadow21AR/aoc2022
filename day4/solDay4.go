package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day4(data *os.File) {
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	re := regexp.MustCompile(`\d+`)
	var res1 int
	var res2 int
	for _, v := range lines {
		ass := re.FindAllStringSubmatch(v, -1)
		temp := make([]int, 0)
		for _, n := range ass {
			x, _ := strconv.Atoi(n[0])
			temp = append(temp, x)
		}
		if (temp[0] >= temp[2] && temp[3] >= temp[1]) || (temp[0] <= temp[2] && temp[3] <= temp[1]) {
			res1++
		}
		if (temp[0] <= temp[3] && temp[3] <= temp[1]) || (temp[2] <= temp[1] && temp[1] <= temp[3]) {
			res2++
		}
	}
	fmt.Println("Answer 4_1 :", res1)
	fmt.Println("Answer 4_2 :", res2)
}

// Answer 4_1 : 496
// Answer 4_2 : 847
