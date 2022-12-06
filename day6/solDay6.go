package day6

import (
	"fmt"
	"os"
)

func Day6() {
	var res, res2 int
	file, err := os.ReadFile("day6/inputDay6.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	for i := 1; i <= len(content)-3; i++ {
		sop := content[i-1 : i+3]
		if check(sop) {
			res = i
			break
		}
	}
	for i := 1; i <= len(content)-13; i++ {
		msg := content[i-1 : i+13]
		if check(msg) {
			res2 = i
			break
		}
	}
	fmt.Println("Answer 6_1: ", res+3, "\nAnswer 6_2: ", res2+13)
}
func check(sop string) bool {
	var a [256]bool
	for _, v := range sop {
		if a[v] {
			return false
		}
		a[v] = true
	}
	return true
}
