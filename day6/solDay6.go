package day6

import (
	"fmt"
	"os"
)

func Day6() {
	file, err := os.ReadFile("day6/inputDay6.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	fmt.Println("Answer 6_1: ", tracker(content, 3)+3, "\nAnswer 6_2: ", tracker(content, 13)+13)
}
func tracker(content string, n int) (out int) {
	for i := 1; i <= len(content)-n; i++ {
		msg := content[i-1 : i+n]
		if check(msg) {
			out = i
			break
		}
	}
	return out
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
